package main

import (
	"fmt"
)

func main() {
	name := "Budi"

	if name == "Eko" {
		fmt.Println("Hallo Eko")
	} else if name == "Joko" {
		fmt.Println("Hallo Joko")
	} else if name == "Budi" {
		fmt.Println("Hallo Budi")
	} else {
		fmt.Println("Hi, kenalan dong")
	}

	if lenght := len(name); lenght > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
