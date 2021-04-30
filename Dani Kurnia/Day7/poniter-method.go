package main

import "fmt"

type Man struct {
	name string
}

func (man *Man) Married() {
	man.name = "Mr. " + man.name
}

func main() {
	eko := Man{"Eko"}
	eko.Married()
	fmt.Println(eko.name)
}
