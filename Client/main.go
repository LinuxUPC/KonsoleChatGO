package main

import (
	"fmt"
	"net"
	"os"
	"bytes"
	"bufio"
	"KonsoleChatGO/ec"
	"KonsoleChatGO/utils"
	"KonsoleChatGO/Client/commands"

)

type Client struct {
	reader *bufio.Reader
	writer *bufio.Writer
	messages chan string
	responses chan string
	uname string
}

func login(conn net.Conn) *Client{

	fmt.Println("type your username")
	fmt.Printf(">")
	var uname string
	fmt.Scan(&uname)
	packet := "login "+uname
	n, err := conn.Write([]byte(packet))
	ec.CheckError(err)
	ec.CheckAllBytesSent(int(n), packet)
	var buffer [512] byte
	n, err = conn.Read(buffer[0:])
	if bytes.Compare(buffer[0:1], []byte("y")) != 0{
		fmt.Fprintf(os.Stderr, "Username already in use")
		//restart
	}
	return &Client{
		reader:bufio.NewReader(conn),
		writer:bufio.NewWriter(conn),
		messages:make(chan string, 999),
		responses:make(chan string, 1),
		uname:uname,
	}
}

func read(c * Client){
	for {
		msg, err := c.reader.ReadString('|')
		if err != nil{
			continue
		}
		if msg[:3] == "msg"{
			c.messages <- msg
		}else{
			c.responses <- msg
		}

	}
}

//handle command  and do network thing

func commandHandler(client *Client){

	go read(client)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(">")
	for scanner.Scan(){

		cmd := scanner.Text()
		comand, _ := utils.ParseCommand(cmd)
		switch comand[0] {
		case  "help":
			commands.Help()
				break
		case "jr":
			err := commands.JoinRoom(comand[1], client.writer, client.responses)
			ec.CheckError(err)

		}
		fmt.Printf(">")

	}

}

func main(){
	fmt.Println(int('\n'))
	fmt.Println("Welcome to konsoleChat")
	fmt.Println("Connecting to server ...")
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:1200")
	ec.CheckError(err)
	fmt.Println("... ...")
	connection, err := net.DialTCP("tcp4", nil, tcpAddr)
	ec.CheckError(err)
	defer connection.Close()
	fmt.Println("Connection successfull!")

	cl := login(connection)

	utils.Cls()
	fmt.Println("success! welcome to KonsoleChat")
	fmt.Println("type \"help\" for help")
	commandHandler(cl)
	end: goto end
}
