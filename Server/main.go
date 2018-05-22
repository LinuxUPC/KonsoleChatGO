package main

import (
	"fmt"
	"net"
	"os"
)

var users = make(map[string]net.Conn)

func clientHandler(conn net.Conn){
	defer conn.Close()
	var buffer [512] byte
	n, err := conn.Read(buffer[0:])
	checkError(err)
	fmt.Println("User", string(buffer[:n]), "logged in.")
	conn.Write([]byte("y")) // don't care about return value
}

func main() {
	port := ":1200"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go clientHandler(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}