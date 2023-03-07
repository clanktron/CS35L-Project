#!/bin/sh

name="${name:="notesbackend"}"
port="${port:="8080"}"

# Build container
docker build -t "$name" .

# Run container
docker run -p "$port":8080 
