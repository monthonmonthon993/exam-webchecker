## Build
FROM golang:1.18.0 AS build-stage

RUN apt-get update && apt-get install -y ca-certificates openssl

ARG cert_location=/usr/local/share/ca-certificates

# Get certificate from "github.com"
RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt
# Get certificate from "proxy.golang.org"
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt
# Update certificates
RUN update-ca-certificates

WORKDIR /app

COPY . /app

RUN go test -v ./...

RUN go mod download

RUN GO111MODULE="on" CGO_ENABLED=0 go build -o /app ./cmd/webchecker-server

## Deploy
FROM alpine:latest

RUN apk add --no-cache bash

WORKDIR /

ENV GIN_MODE=release

COPY --from=build-stage /app/webchecker-server /webchecker-server

EXPOSE 8080

ENTRYPOINT ["./webchecker-server"]