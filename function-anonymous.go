package main

import "fmt"

type Blacklist func(string) bool

func blacklistUser(name string, blacklist Blacklist) string {
	if blacklist(name) {
		return "You Are Blocked " + name
	} else {
		return name
	}
}

func main() {

	fmt.Println(blacklistUser("kresna", func(name string) bool {
		return name == "kresna"
	}))

}
