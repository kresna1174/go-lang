package main

import "fmt"

func variableFunc(names string) string {

	return "Hello " + names

}

func main() {
	funcVariable := variableFunc

	fmt.Println(funcVariable("Kresna"))
}
