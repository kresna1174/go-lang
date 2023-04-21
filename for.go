package main

import "fmt"

func main() {
	array_list := []string{
		"Kresna",
		"Cahyono",
		"Kresna 1",
		"Cahyono 1",
	}

	for i, name := range array_list {
		fmt.Println("Index", i, "=", name)
	}

}