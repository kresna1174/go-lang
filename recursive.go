package main

import "fmt"

func factorial(number int) int {
	result := 1
	for i := number; i > 0; i-- {
		result *= i
	}
	return result
}

func recursiveFactorial(number int) int {
	if number == 1{
		return 1
	} else {
		return number * recursiveFactorial(number-1);
	}
}

func main() {

	fmt.Println(factorial(5))
	fmt.Println(recursiveFactorial(5))

}