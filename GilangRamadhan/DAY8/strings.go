package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("gilang ramadhan", "freya"))
	fmt.Println(strings.Contains("gilang ramadhan", "gilang"))

	fmt.Println(strings.Split("gilang ramadhan", " "))

	fmt.Println(strings.ToLower("Gilang Ramadhan"))
	fmt.Println(strings.ToUpper("Gilang Ramadhan"))
	fmt.Println(strings.ToTitle("Gilang Ramadhan"))

	fmt.Println(strings.Trim("    gilang ramadhan    ", " "))
	fmt.Println(strings.ReplaceAll("gilang gilang gilang", "gilang", "freya"))
}
