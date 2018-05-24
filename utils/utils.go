package utils

import (
	"os/exec"
	"os"
)

func Cls(){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

/*func ParseCommand(cmd string) [] string{
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