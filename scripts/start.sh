#!/bin/bash

function start_server() {
    CERT_FILE="certs/localhost.pem" \
    KEY_FILE="certs/localhost-key.pem" \
    SERVICE_ADDR=":8080" \
    go run main.go
}
