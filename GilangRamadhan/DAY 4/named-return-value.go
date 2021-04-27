package main

import "fmt"

func getcompletename() (firstname string, lastname string) {
	firstname = "gilang"
	lastname = "ramadhan"

	return
}

func main() {
	a, b := getcompletename()
	fmt.Println(a)
	fmt.Println(b)
}
