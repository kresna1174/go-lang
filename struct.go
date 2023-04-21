package main

import "fmt"

type Customer struct {
	Name, FirstName, LastName string
	Age                       int
}

func main() {

	var customer Customer
	customer.Name = "Kresna"
	customer.FirstName = "Cahyono"
	customer.LastName = "Cahyono"
	customer.Age = 20

	fmt.Println(customer)

	customer2 := Customer {
		Name: "Kresna",
		FirstName: "Kresna",
		LastName: "Cahyono",
		Age: 20,
	}

	
	fmt.Println(customer2)
	customer3 := Customer{"Kresna", "Bambang", "Yudhoyono", 20}
	fmt.Println(customer3.Name)

}