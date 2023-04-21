package main

import "fmt"

func SayString() interface{} {
	return map[string]string{
		"name": "Kresna",
	}
}

func main() {

	result := SayString()

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	case bool:
		fmt.Println("Boolean", value)
	default:
		fmt.Println(value)

	}

}
