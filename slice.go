package main

import "fmt"

func main() {

	array := [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	result_array := array[0:12]

	fmt.Println(result_array);

	fmt.Println(len(array))
	result_array[0] = "januari now"

	fmt.Println(result_array)

	as := make(map[int]string);
	
	for counter := 1; counter <= 10; counter++ {
	}

	fmt.Println(as);

}