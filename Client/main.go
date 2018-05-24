package main

import (
	"fmt"
	"net"
	"os"
	"bytes"
	"os/exec"
	"bufio"
)

func cls(){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func checkAllBytesSent(n int, intended_send string){
	if n != len(intended_send){
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

/*func parseCommand(cmd string) [] string{
	word :=""
	reading := false
	var words []string
	for _, char := range cmd{
		if char == ' '{
			if word != ""{
				words = append(words, word)
				word = ""
			}
		}else if char == '"' || reading{
			if()
			reading = true
			word += string(char)

		}

		if char != ' ' && char != '"'{
			word += string(char)
		}
	}
}*/

//handle command  and do network thing

func commandHandler(conn net.Conn, uname string ){
	defer conn.Close()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		cmd := scanner.Text()
		//command := parseCommand(cmd)
	}

}

func main(){
	fmt.Println(int('\n'))
	fmt.Println("Welcome to konsoleChat")
	fmt.Println("Connecting to server ...")
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:1200")
	checkError(err)
	fmt.Println("... ...")
	connection, err := net.DialTCP("tcp4", nil, tcpAddr)
	checkError(err)
	fmt.Println("Connection successfull!")

login:
	fmt.Println("type your username")
	var uname string
	fmt.Scan(&uname)
	n, err := connection.Write([]byte(uname))
	checkError(err)
	checkAllBytesSent(int(n), uname)
	var buffer [512] byte
	n, err = connection.Read(buffer[0:])
	if bytes.Compare(buffer[0:1], []byte("y")) != 0{
		fmt.Fprintf(os.Stderr, "Username already in use")
		goto login
	}
	fmt.Println("success!")
	commandHandler(connection, uname)





	end: goto end
}
