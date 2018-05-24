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
	username := string(buffer[:n])
	fmt.Println("User", username, "logged in.")
	if users[username] == nil {
		users[username] = conn
	}
	for {
		_, err := conn.Read(buffer[0:])
		checkError(err)
		/**recived = split(buffer, n)**/
	}
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
