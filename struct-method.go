package main

import "fmt"

type Customer struct {
	Name, FirstName, LastName string
	Age                       int
}

func (customer Customer) sayCustomer(name string) {
	fmt.Println("Hello", name, " my name is", customer.Name)
}

func main() {

	var variable Customer
	variable.Name = "Kresna"
	// variable.Name = "Kresna"
	variable.sayCustomer("Cahyono")

}
