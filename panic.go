package main

import "fmt"

func endApp() {
	fmt.Println("Stop")

	msg := recover()
	if msg != nil {
		fmt.Println("Terjadi Error:", msg)
	}
}

func startApp(boolean bool) {
	defer endApp()
	if boolean {
		panic("AKU PANIK")
	}
	fmt.Println("Start App")

}

func main() {

	startApp(false)

}
