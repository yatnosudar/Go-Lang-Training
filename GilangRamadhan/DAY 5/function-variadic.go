package main

import "fmt"

func sumALL(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumALL(10, 20, 50, 20, 100, 69)
	fmt.Println(total)

	slice := []int{10, 20, 20, 50}
	total = sumALL(slice...)
	fmt.Println(total)
}
