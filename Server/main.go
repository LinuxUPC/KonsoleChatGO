package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"KonsoleChatGO/ec"
	"KonsoleChatGO/utils"
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

func joinRoom(me * Client, all *[]Client, roomname string) {
	roomExists := false
	for _, current := range *all {
		if &current != me{
			if current.room == roomname {
				roomExists = true
				break
			}
		}
	}

	if roomExists {
		me.room = roomname
		me.writer.Write([]byte("jr_success|"))
	}else {
		me.writer.Write([]byte("jr_fail room_exists|"))
	}
}

func leaveRoom(client *Client) {
	client.room = ""
	client.writer.Write([]byte("lr_success|"))
}

func login(me *Client, all *[]Client, alias string) {
	aliasExists := false
	for _, current := range *all {
		if &current != me{
			if current.room == alias {
				aliasExists = true
				break
			}
		}
	}

	if aliasExists{
		me.writer.Write([]byte("li_fail alias_exists|"))
	}else{
		me.alias = alias
		me.writer.Write([]byte("li_success|"))
	}
}

func logout(client *Client) {
	client.alias = ""
	client.writer.Write([]byte("lo_success|"))
}

func send(me *Client, all *[]Client, msg string) {
	for _, current := range *all {
		if current.room != me.room {
			current.writer.Write([]byte("msg " + msg + "|"))
		}
	}
}

func clientConnection(me * Client, all *[]Client){
	defer me.conn.Close()
	for {
		msg, err := me.reader.ReadString('|')
		ec.CheckError(err)
		comands ,err := utils.ParseCommand(msg)
		ec.CheckError(err)
		switch comands[0]{
		case "cr":{

		}
		case "jr":{
			if len(comands) != 2{
				me.writer.Write([]byte("wf|"))
			}else{
				joinRoom(me, all, comands[1])
			}
		}
		case "lr":{
			if len(comands) != 1{
				me.writer.Write([]byte("wf|"))
			}else{
				leaveRoom(me)
			}
		}
		case "li":{
			if len(comands) != 2{

			}else{
				login(me, all, comands[1])
			}
		}
		case "lo":{
			if len(comands) != 1{

			}else{
				logout(me)
			}
		}
		case "msg":{
			if len(comands) != 2{

			}else{
				send(me, all, comands[1])
			}
		}
		default:{
			me.writer.Write([]byte("wc|"))
		}
		}
	}
}

var clients []Client
var newC * Client

func main() {
	port := ":1200"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		fmt.Println("Hola")
		conn, err := listener.Accept()
		if err == nil {
			continue
		}
		makeClient(newC, conn)
		clients = append(clients, *newC)
		fmt.Println("hola")
		go clientConnection(&(clients[len(clients)-1]), &clients)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
