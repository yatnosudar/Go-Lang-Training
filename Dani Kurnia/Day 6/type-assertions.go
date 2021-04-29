package main

import "fmt"

func random() interface{} {
	return "Ups"
}

func main() {
	var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case int:
		fmt.Println("value", value, "is int")
	case string:
		fmt.Println("value", value, "is string")
	default:
		fmt.Println("Unknown")
	}

}
