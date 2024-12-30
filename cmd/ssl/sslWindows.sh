#!/bin/bash

set -e # Exit on error
set -x # Debugging output

# Variables
SERVER_CN=localhost
CA_SUBJECT="//CN=ca"
SERVER_SUBJECT="//CN=${SERVER_CN}"

# Step 1: Generate Certificate Authority + Trust Certificate (ca.crt)
openssl genrsa -passout pass:1111 -des3 -out ca.key 4096
openssl req -passin pass:1111 -new -x509 -sha256 -days 365 -key ca.key -out ca.crt -subj "${CA_SUBJECT}"

# Step 2: Generate the Server Private Key (server.key)
openssl genrsa -passout pass:1111 -des3 -out server.key 4096

# Step 3: Get a certificate signing request from the CA (server.csr)
openssl req -passin pass:1111 -new -key server.key -out server.csr -subj "${SERVER_SUBJECT}"

# Step 4: Create extfile.cnf for subjectAltName
echo "subjectAltName=DNS:${SERVER_CN},IP:0.0.0.0" > extfile.cnf

# Step 5: Sign the certificate with the CA we created (self-signing) - server.crt
openssl x509 -req -extfile extfile.cnf -passin pass:1111 -sha256 -days 1095 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt

# Step 6: Convert the server certificate to .pem format (server.pem) - usable by gRPC
openssl pkcs8 -topk8 -nocrypt -passin pass:1111 -in server.key -out server.pem