package main

import "fmt"

func InterfaceKosong(number int) interface{} {
	if number == 1 {
		return 1
	} else if number == 2 {
		return true
	} else {
		return "String Interface Kosong"
	}
}

func main() {

	var kosong interface{} = InterfaceKosong(3)
	fmt.Println(kosong)

}
