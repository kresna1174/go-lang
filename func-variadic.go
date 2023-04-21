package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func fullName() (firstname, lastname, callme string) {
	firstname = "Kresna";
	lastname = "Cahyono";
	callme = "Nana";

	return
}

func main() {
	fmt.Println(sumAll(10,20,30,40,50))
	fmt.Println(fullName())
}