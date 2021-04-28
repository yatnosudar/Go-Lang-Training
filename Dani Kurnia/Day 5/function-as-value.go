package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodBye := getGoodBye
	result := goodBye("Eko")
	fmt.Println(result)
	fmt.Println(getGoodBye("Dadan"))
}
