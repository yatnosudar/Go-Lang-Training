package main

import "fmt"

func main() {
	counter := 0
	name := "gilang"

	increment := func() {
		name := "freya" //  (:) < pake ini supaya bisa ke baca 2
		fmt.Println("increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
