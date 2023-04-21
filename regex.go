package main

import (
	"fmt"
	"regexp"
)

func main() {

	regex := regexp.MustCompile(`k([a-z])a`)

	fmt.Println(regex.MatchString("kresna"))
	fmt.Println(regex.MatchString("kia"))
	fmt.Println(regex.MatchString("kana"))
	fmt.Println(regex.MatchString("kaca"))
	fmt.Println(regex.MatchString("kaba"))

	fmt.Println(regex.FindAllString("kia kaa kaa kam mu kasa", -1))

}