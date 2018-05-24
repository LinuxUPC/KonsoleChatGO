package main

import (
	"fmt"
	"net"
	"os"
	"bytes"
	"bufio"
	"KonsoleChatGO/ec"
	"KonsoleChatGO/utils"
)

//handle command  and do network thing

func commandHandler(conn net.Conn, uname string ){
	defer conn.Close()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(">")
	for scanner.Scan(){

		cmd := scanner.Text()
		comand, _ := utils.ParseCommand(cmd)
		fmt.Println("The contents of the command are:")
		for _, strin := range comand{
			fmt.Println(strin)
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
