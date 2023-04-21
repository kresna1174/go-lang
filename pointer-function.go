package main

import "fmt"

type Address struct {
	Kota, Provinsi, Negara string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Negara = "Indonesia"
}

func main() {

	alamat := Address{"New London", "West London", "Inggris"}

	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)

}
