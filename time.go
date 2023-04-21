package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	parse, _ := time.Parse("2006-01-02", "2023-01-01")
	fmt.Println(now)
	fmt.Println(parse)
}
