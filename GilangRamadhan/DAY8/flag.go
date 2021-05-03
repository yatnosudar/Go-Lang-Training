package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your host")
	var user *string = flag.String("user", "root", "Put your database user")
	var password *string = flag.String("password", "root", "Put your database password")
	var number *int = flag.Int("number", 100, "put yout number")

	flag.Parse()

	fmt.Println("host :", *host)
	fmt.Println("user :", *user)
	fmt.Println("password :", *password)
	fmt.Println("Number :", *number)
}
