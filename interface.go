package main

import "fmt"

type HasName interface {
	GetName() string
	GetAge() string
}

type Person struct {
	Name string
	Age int
}

func (person Person) GetName() {
	fmt.Println("Hii Person From", person.Name)
}

func (person Person) GetAge() {
	fmt.Println(person.Name, "is", person.Age, "Years Old")
}

type Animal struct {
	Name string
	Age string
}

func (animal Animal) GetName() string {
	return "Hi Animal " + animal.Name
}

func (animal Animal) GetAge() string {
	return animal.Name + " is " + animal.Age +" Years Old"
}

func main() {

	var person Person
	person.Name = "Kresna"
	person.Age = 20
	person.GetName()
	person.GetAge()

	animal := Animal{
		Name: "Kucing",
		Age: "2",
	}

	fmt.Println(animal.GetName())
	fmt.Println(animal.GetAge())

}
