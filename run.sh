#!/bin/sh
export projectPATH="$GOPATH/src/konsoleCHat"

go build -o "$projectPATH/tmp/server" ./Server
go build -o "$projectPATH/tmp/client" ./Client

gnome-terminal --tab --title="Server"  -- $projectPATH/tmp/server
gnome-terminal --tab --title="Cleint1" -- $projectPATH/tmp/client
gnome-terminal --tab --title="Client2" -- $projectPATH/tmp/client

unset projectPATH

