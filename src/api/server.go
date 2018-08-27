package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hammi85/swerve/src/db"
)

// NewServer creates a new instance
func NewServer(listener string, dynDB *db.DynamoDB) *Server {
	server := &Server{
		Listener: listener,
		db:       dynDB,
	}

	server.Server = &http.Server{
		Addr:    listener,
		Handler: server.getMux(),
	}

	return server
}

// Listen to socket
func (s *Server) Listen() error {
	log.Printf("API listening to %s", s.Listener)
	return s.Server.ListenAndServe()
}

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", 405)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"ok\"}"))
}

func (s *Server) purgeDomain(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", 405)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"ok\"}"))
}

func (s *Server) fetchAllDomains(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", 405)
		return
	}

	s.db.FetchAll()
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"ok\"}"))
}

func (s *Server) registerDomain(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", 405)
		return
	}

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	var domain db.Domain

	err := json.NewDecoder(r.Body).Decode(&domain)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = s.db.InsertDomain(domain)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"ok\"}"))
}

// GetMux returns the router handler
func (s *Server) getMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/domain", s.fetchAllDomains)
	mux.HandleFunc("/domain/register", s.registerDomain)
	mux.HandleFunc("/domain/purge", s.purgeDomain)
	mux.HandleFunc("/health", s.health)

	return mux
}
