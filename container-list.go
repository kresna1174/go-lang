package main

import (
	"container/list"
	"fmt"
)

func main() {

	data := list.New()
	data.PushBack("Kresna")
	data.PushBack("Cahyono")
	data.PushBack("Wijoyo")
	data.PushBack("Kusumo")

	// dari depan
	for i := data.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	// dari belakang
	for i := data.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}

}
