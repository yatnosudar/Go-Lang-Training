package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	// person := NewMap("eko")
	// fmt.Println(person)

	var person map[string]string = NewMap("gilang")

	if person == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(person)
	}
}
