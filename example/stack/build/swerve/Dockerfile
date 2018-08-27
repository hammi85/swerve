# BUILD
FROM golang:latest as build

LABEL maintainer="sejamich@googlemail.com"

WORKDIR /go/src/swerve
COPY . .

#RUN go get -d -v ./...
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cataprox -v -ldflags "-extldflags '-static'" -a -installsuffix cgo main.go

# RUNTIME
#FROM scratch

#MAINTAINER Jan Michalowsky <sejamich@googlemail.com>

#COPY --from=build /go/src/cache/cache /

EXPOSE 80

RUN cp swerve /swerve

CMD [ "/swerve" ]