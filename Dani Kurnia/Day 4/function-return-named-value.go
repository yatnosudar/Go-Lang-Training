package main

import "fmt"

func getData() (firstName, lastName string, age int) {
	firstName = "Eko"
	lastName = "Kurniawan"
	age = 22

	return
}

func main() {
	a, b, c := getData()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
