# BUILD
FROM golang:latest as build

LABEL maintainer="sejamich@googlemail.com"

WORKDIR /go/src/github.com/hammi85/swerve
COPY . .

RUN echo $GOPATH
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o swerve -v -ldflags "-extldflags '-static'" -a -installsuffix cgo main.go

# RUNTIME
FROM scratch

MAINTAINER Jan Michalowsky <sejamich@googlemail.com>

COPY --from=build /go/src/github.com/hammi85/swerve/swerve /swerve

CMD [ "/swerve" ]