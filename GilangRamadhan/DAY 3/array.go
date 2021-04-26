package main 

import "fmt"

func main ( ) {

	var names [3]string

	names[0] = "gilang"
	names[1] = "Ramadhan"
	names[2] = "freya"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names)

	var values = [3]int {
		98,
		90,
		100,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10] string
	  
	fmt.Println(len(lagi))
}