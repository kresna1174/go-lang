package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("HOST", "LOCALHOST", "YOUR HOST")
	username := flag.String("USERNAME", "ROOT", "YOUR USERNAME")
	password := flag.String("PASSOWRD", "ROOT", "YOUR PASSOWRD")

	flag.Parse()

	fmt.Println(*host, *username, *password)
}