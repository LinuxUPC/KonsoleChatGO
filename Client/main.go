package main

import (
	"fmt"
	"net"
	"os"
	"bytes"
)

func checkAllBytesSent(n int, uname string){
	if n != len(uname){
		fmt.Fprintf(os.Stderr, "Fatal error: could not send all data to server")
		os.Exit(1)
	}
}

func checkError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main(){
	fmt.Println("Welcome to konsoleChat")
	fmt.Println("Connecting to server ...")
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:1201")
	checkError(err)
	fmt.Println("... ...")
	connection, err := net.DialTCP("tcp4", nil, tcpAddr)
	checkError(err)
	fmt.Println("Connection successfull!")

login:
	fmt.Println("type your username")
	var uname string
	fmt.Scanln(uname)
	n, err := connection.Write([]byte(uname))
	checkError(err)
	checkAllBytesSent(int(n), uname)
	var buffer [512] byte
	n, err = connection.Read(buffer[0:])
	if bytes.Compare(buffer[0:1], []byte("y")) != 0{
		fmt.Fprintf(os.Stderr, "Username already in use")
		goto login
	}
}