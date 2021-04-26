package main

import (
	"fmt"
)

func main() {
	person := map[string]string{
		"name":    "Eko",
		"address": "Subang",
	}

	fmt.Println(person)
	fmt.Println("name")
	fmt.Println("address")

	person["title"] = "Programmer"
	fmt.Println(person)

	book := make(map[string]string)

	book["title"] = "Belajar Go-Lang"
	book["author"] = "Eko"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}
