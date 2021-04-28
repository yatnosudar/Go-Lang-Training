package main

import "fmt"

type blacklist func(string) bool

func registerUser(name string, blacklisit blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"

	}

	registerUser("admin", blacklist)
	registerUser("gilang", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})

	registerUser("gilang", func(name string) bool {
		return name == "root"
	})

}
