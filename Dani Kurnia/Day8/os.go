package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument :")
	fmt.Println(args)
	fmt.Println(args[1])
	fmt.Println(args[2])

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", name)
	} else {
		fmt.Println("Error", err.Error())

		username := os.Getenv("App_Username")
		password := os.Getenv("App_Password")
		fmt.Println(username)
		fmt.Println(password)
	}
}
