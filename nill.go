package main

import "fmt"

func mapping(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {

	var person map[string]string = mapping("Kresna")
	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}

	var me = map[string]string{
		"name": "Kresna",
	}

	fmt.Println(me["name"])

}
