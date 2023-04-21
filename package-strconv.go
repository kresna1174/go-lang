package main

import (
	"fmt"
	"strconv"
)

func main() {

	boolean, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	number, err := strconv.ParseInt("10000", 10, 32)

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	float, err := strconv.ParseFloat("10,000", 32)

	if err == nil {
		fmt.Println(float)
	} else {
		fmt.Println(err.Error())
	}

}
