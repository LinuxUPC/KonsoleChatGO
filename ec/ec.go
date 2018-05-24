package ec

import (
	"fmt"
	"os"
)

func CheckAllBytesSent(n int, intended_send string){
	if n != len(intended_send){
		fmt.Fprintf(os.Stderr, "Fatal error: could not send all data to server")
		os.Exit(1)
	}
}

func CheckError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}


