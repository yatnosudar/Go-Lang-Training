package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	gilang := Man{"Gilang"}
	gilang.Married()

	fmt.Println(gilang.Name)
}
