#!/bin/sh
#export projectPATH="$GOPATH/src/KonsoleChatGO"


go build -o "$projectPATH/tmp/server" ./Server
go build -o "$projectPATH/tmp/client" ./Client

runAll(){
    gnome-terminal --tab --title="Server" --geometry=80x50+200+10 -- $projectPATH/tmp/server;
    gnome-terminal --tab --title="Cleint1" --geometry=80x25+950+10 -- $projectPATH/tmp/client;
    gnome-terminal --tab --title="Client2" --geometry=80x25+950+500 -- $projectPATH/tmp/client;
}

runAll


unset runAll
unset projectPATH

