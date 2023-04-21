package main

import "fmt"

type Address struct {
	Kota, Provinsi, Negara string
}

func main() {

	address1 := Address{"Tulungagung", "Jawa Timur", "Indonesia"}
	address2 := &address1

	address2.Kota = "Nganjuk"

	*address2 = Address{"Ngantru", "Jatim", "Indo"}
	fmt.Println(address1)
	fmt.Println(address2)

	address4 := new(Address)
	address4.Kota = "Ngunut"
	fmt.Println(address4)

}