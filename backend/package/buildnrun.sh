#!/bin/sh

username="${username:=clanktron}"
name="${name:="notesbackend"}"
fullname="$username/$name"
port="${port:="4000"}"

# Build container
docker build  -t "$fullname" -f package/dockerfile . 

# Run container
docker run --name="$name" -p "$port":"$port" "$fullname"
