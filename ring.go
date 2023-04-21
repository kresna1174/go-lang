package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {

	var ring *ring.Ring = ring.New(5)

	for i := 0; i < ring.Len(); i++ {
		ring.Value = "Ini Ring Ke " + strconv.FormatInt(int64(i), 10)
		ring = ring.Next()
	}

	ring.Do(func(value any) {
		fmt.Println(value)
	})

}
