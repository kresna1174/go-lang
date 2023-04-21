package main

import (
	"errors"
	"fmt"
)

func cekNilai(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {

	nilai, err := cekNilai(100, 0)
	if err == nil {
		fmt.Println(nilai)
	} else {
		fmt.Println("Error", err.Error())
	}

}
