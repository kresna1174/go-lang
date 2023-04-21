package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) string {
	filtered := filtered

	return filtered(name)
}

func filtered(name string) string {

	switch name {
	case "cok":
		return ".*."
	case "anjing":
		return ".*."
	case "bangsat":
		return ".*."
	case "babi":
		return ".*."
	default:
		return name
	}

}

func main() {

	fmt.Println(sayHelloWithFilter("cok", filtered))
	fmt.Println(sayHelloWithFilter("anjing", filtered))
	fmt.Println(sayHelloWithFilter("bangsat", filtered))
	fmt.Println(sayHelloWithFilter("babi", filtered))
	fmt.Println(sayHelloWithFilter("kresna ganteng", filtered))

}
