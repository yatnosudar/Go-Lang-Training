package main

import "fmt"

type filter func(string) string

func sayHelloWithFilter(name string, filter filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "Kasar"
	} else {
		return name
	}
}

func parameter() {
	sayHelloWithFilter("GILANG", spamFilter)
	sayHelloWithFilter("anjing", spamFilter)
}
