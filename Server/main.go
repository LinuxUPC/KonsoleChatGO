package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)

type Client struct {
	conn net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
	messages chan string
	responses chan string
	alias string
	room string
}

func makeClient(new * Client, conn net.Conn){
	new = &Client{
		conn:conn,
		reader:bufio.NewReader(conn),
		writer:bufio.NewWriter(conn),
		messages:make(chan string),
		responses:make(chan string),
		alias:"",
		room:"",
	}
}

func split(str string, separator byte){

}

func handleMessage(message string, conn net.Conn){

}

func clientConnection(me * Client, all []Client){
	defer me.conn.Close()
	for {
		me.reader.ReadBytes('|')
	}
}

func main() {
	port := ":1200"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	var clients []Client

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		var new Client
		makeClient(&new, conn)
		clients = append(clients, new)
		go clientConnection(&new, &clients)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
