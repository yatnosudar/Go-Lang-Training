package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hi", name, "my name is", customer.Name)
}

func (a Customer) sayHuu() {
	fmt.Println("Huu from", a.Name)
}

func main() {
	var eko Customer

	eko.Name = "Eko"
	eko.Address = "Indonesia"
	eko.Age = 30

	eko.sayHi("Budi")
	eko.sayHuu()

	// fmt.Println(eko.Name)
	// fmt.Println(eko.Address)
	// fmt.Println(eko.Age)
	// fmt.Println(eko)

	// joko := Customer{
	// 	Name:    "Joko",
	// 	Address: "Surabaya",
	// 	Age:     28,
	// }
	// fmt.Println(joko)

	// budi := Customer{"Budi", "Bandung", 32}
	// fmt.Println(budi)
}
