package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke", counter)
		counter++
	}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke", counter)
	}

	slice := []string{"gilang", "rama", "freya", "allan"}
	for i := 0; 1 < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index", i, "=" value)
	}

	person := make (map[string]string)
	person["name"] = "gilang"
	person["title"] = "rebahan"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}
