package main

import "fmt"

func main() {

	name := "gilang"

	switch name {
	case "gilang":
		fmt.Println("hello gilang")
		fmt.Println("hello gilang")
	case "rama":
		fmt.Println("hello rama")
		fmt.Println("hello rama")
	default:
		fmt.Println("boleh kenalan")
		fmt.Println("boleh kenalan")

	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}

}
