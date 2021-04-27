package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hae"
	} else {
		return "Halo" + name
	}
}

func main() {
	result := getHello("Eko")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
