package main

import (
	"fmt"
)

func main() {
	name := "EkoEko"

	switch name {
	case "Eko":
		fmt.Println("Hallo Eko")
		fmt.Println("Hallo Eko")
	case "Budi":
		fmt.Println("Hallo Budi")
	default:
		fmt.Println("Hai, boleh kenalan?")
	}

	switch lenght := len(name); lenght > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah cukup")
	}
}
