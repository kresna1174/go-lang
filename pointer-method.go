package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) married() {
	man.Name = "Mr . " + man.Name
}

func main() {

	YourName := Man{"Kresna"}

	YourName.married()
	fmt.Println(YourName.Name)

}
