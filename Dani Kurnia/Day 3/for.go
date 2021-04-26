package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}
	slice := []string{"Eko", "Kurniawan", "Khannedy"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)

	person["Nama"] = "Joko"
	person["Title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
