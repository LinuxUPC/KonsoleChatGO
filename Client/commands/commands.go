package commands

import (
	"fmt"
	"bufio"
	"errors"
)

func Help(){
	fmt.Println("Usage: jr \"room name", "  To join a room")
	fmt.Println("Usage: cr \"room name", "  To create a room")
}

func JoinRoom(room string, netWrite *bufio.Writer, resp chan string) error {
	netWrite.Write(([]byte("jr " + room)))
	for{
		select{
		case x, ok := <-resp:
			if(ok){
				if x == "ok"{
					return nil
				}else{
					return errors.New("Room does not exist!")
				}
			}else{
				return errors.New("channel closed!")
			}
		default:
			break

		}
	}
}