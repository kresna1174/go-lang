package main

import (
	"fmt"
	"os"
)

func main() {
	Argusment := os.Args

	if len(Argusment) > 1 {
		for i := 0; i < len(Argusment); i++ {
			fmt.Println(Argusment[i])
		}
	} else {
		fmt.Println(Argusment)

	}

	hostname, err := os.Hostname()

	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err)
	}
}
