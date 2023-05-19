#!/bin/bash

mkdir -p http2/certs

cd http2/certs

# Generate a self-signed certificate for the server
openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout server.key -out server.crt
