package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "Eko"
	sayHelloTo(firstName, "Khannedy")
	sayHelloTo("Budi", "Anduk")
}
