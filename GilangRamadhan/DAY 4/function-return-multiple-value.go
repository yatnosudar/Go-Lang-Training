package main

import "fmt"

func getFullName() (string, string) {
	return "gilang", "ramadhan"
}

func main() {
	firstname, lastname := getFullName()
	fmt.Println(firstname)
	fmt.Println(lastname)

}
