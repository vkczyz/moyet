#!/bin/sh
openssl genrsa -out server.key 2048
openssl req -x509 -new -nodes -key server.key -sha256 -out server.crt
