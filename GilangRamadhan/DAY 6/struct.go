package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHI(name string) {
	fmt.Println("Hello", name, "My Name Is", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("HUUUUUU FROM", a.Name)
}

func main() {
	var gilang Customer
	gilang.Name = "gilang"
	gilang.Address = "indonesia"
	gilang.Age = 17

	gilang.sayHI("BUDI")
	gilang.sayHuuu()

	fmt.Println(gilang.Name)
	fmt.Println(gilang.Address)
	fmt.Println(gilang.Age)

	freya := Customer{
		Name:    "freya",
		Address: "USA",
		Age:     17,
	}
	fmt.Println(freya)

	rama := Customer{"rama", "bandung", 18}
	fmt.Println(rama)

}
