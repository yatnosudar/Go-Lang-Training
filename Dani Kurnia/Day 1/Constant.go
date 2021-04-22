package main

import "fmt"

func main() {
	const firstName string = "Eko"
	const lastName = "Kurniawan"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	const (
		country = "Indonesia"
		city    = "Bandung"
	)
	fmt.Println(country)
	fmt.Println(city)
}
