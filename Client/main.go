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



func read(reader *bufio.Reader, mesages chan <-string, responses chan <- string){
	for {
		msg, err := reader.ReadString('|')
		ec.CheckError(err)
		mesages <- msg
		if msg[:3] == "msg"{
			mesages <- msg
		}else{
			responses <- msg
		}

	}
}

//handle command  and do network thing

func commandHandler(conn net.Conn, uname string ){
	defer conn.Close()
	netReader := bufio.NewReader(conn)
	netWriter := bufio.NewWriter(conn)
	mesages := make(chan string, 1000)
	response := make(chan string, 1)
	go read(netReader,mesages, response)
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
			err := commands.JoinRoom(comand[1], netWriter, response)
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
	fmt.Println("Connection successfull!")

login:
	fmt.Println("type your username")
	fmt.Printf(">")
	var uname string
	fmt.Scan(&uname)
	n, err := connection.Write([]byte(uname))
	ec.CheckError(err)
	ec.CheckAllBytesSent(int(n), uname)
	var buffer [512] byte
	n, err = connection.Read(buffer[0:])
	if bytes.Compare(buffer[0:1], []byte("y")) != 0{
		fmt.Fprintf(os.Stderr, "Username already in use")
		goto login
	}
	utils.Cls()
	fmt.Println("success! welcome to KonsoleChat")
	fmt.Println("type \"help\" for help")


	commandHandler(connection, uname)





	end: goto end
}
