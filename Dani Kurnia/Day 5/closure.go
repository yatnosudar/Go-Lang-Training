package main

import "fmt"

func main() {
	name := "Eko"
	counter := 0

	increment := func() {
		name := "Budi"
		fmt.Println(name)
		fmt.Println("Increment")
		counter++
	}
	increment()
	increment()

	fmt.Println(name)
	fmt.Println(counter)
}
