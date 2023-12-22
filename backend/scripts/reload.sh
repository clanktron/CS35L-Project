#!/bin/sh
# Hot reload
find . -name \*.go | entr -cr go run .
# git ls-files | grep .go | entr go run main.go
