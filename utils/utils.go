package utils

import (
	"os/exec"
	"os"
	"errors"
)

func Cls(){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ParseCommand(cmd string) ([] string,error){
	word :=""
	reading := false
	var words []string
	for _, char := range cmd{
		if char == '"' || reading{
			if reading && char == '"'{
				reading = false
				words = append(words, word)
				word = ""

			}else if !reading{
				if word != ""{
					return make([]string,1), errors.New("Bad string")
				}
				reading = true
			}else{
				word += string(char)
			}


		}else if char == ' '{
			if word != ""{
				words = append(words, word)
				word = ""
			}
		}else{
			word += string(char)
		}

	}
	return words, nil
}