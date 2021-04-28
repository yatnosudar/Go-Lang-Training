package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blackList Blacklist) {
	if blackList(name) {
		fmt.Println("You are blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

// func blackListAdmin(name string) bool {
// 	return name == "Admin"
// }

// func blackListRoot(name string) bool {
// 	return name == "Root"
// }

func main() {
	blackList := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", blackList)
	registerUser("eko", blackList)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("eko", func(name string) bool {
		return name == "eko"
	})
}
