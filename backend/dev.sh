#!/bin/sh
# Barebones golang dev container
#
# Usage: run this followed by any other "docker run" flags.
#
Container="${Container:=clanktron/godev}"
Port="${Port:="4000"}"
ContainerName="${ContainerName:=develop}"
# 
alias docker=nerdctl
#
docker run -p "$Port":"$Port" -it -v "$PWD":/code "$@" --name develop "$Container"
