package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Gilang",
		"address": "Bandung",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	var book map[string]string = make(map[string]string)
	book["book"] = "Belajar Go-Lang"
	book["author"] = "Gilang"
	book["ups"] = "salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
