package main

import (
	"fmt"
	"belajar-go/database"
)

func main() {
	database := database.GetConn()
	fmt.Println(database)
}