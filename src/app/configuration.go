package app

import (
	"os"
	"strings"
)

const (
	envPrefix = "SWERVE_"
)

// getOSPrefixEnv get os env
func getOSPrefixEnv(s string) *string {
	if e := strings.TrimSpace(os.Getenv(envPrefix + s)); len(e) > 0 {
		return &e
	}

	return nil
}

// FromEnv read the config from envs
func (c *Configuration) FromEnv() {
	if api := getOSPrefixEnv("API"); api != nil {
		c.APIListener = *api
	}

	if httpListener := getOSPrefixEnv("HTTP"); httpListener != nil {
		c.HTTPListener = *httpListener
	}

	if httpsListener := getOSPrefixEnv("HTTPS"); httpsListener != nil {
		c.HTTPSListener = *httpsListener
	}
}

// FromParameter read config from application parameter
func (c *Configuration) FromParameter() {
}

// NewConfiguration creates a new instance
func NewConfiguration() *Configuration {
	return &Configuration{
		HTTPListener:  ":8080",
		HTTPSListener: ":8081",
		APIListener:   ":8082",
	}
}
