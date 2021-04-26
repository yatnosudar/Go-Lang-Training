package main

import "fmt"

func main() {
	var name = "ramadhan"

	if name == "gilang" {
		fmt.Println("hello gilang")
	} else if name == "rama" {
		fmt.Println("hello rama")
	} else if name == "freya" {
		fmt.Println("hello freya")
	} else {
		fmt.Println("hi, kenalan dong")
	}

	if length := len(name); length > 5 {
		fmt.Println("terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
