package main

import "fmt"

func sayHelloTo(firstname string, lastname string) {
	fmt.Println("hello", firstname, lastname)
}

func main() {
	firstname := "gilang"
	sayHelloTo(firstname, "gilang")
	sayHelloTo("gilang", "freya")
}
