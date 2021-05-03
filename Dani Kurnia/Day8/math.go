package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.8))
	fmt.Println(math.Round(1.2))
	fmt.Println(math.Floor(1.7))
	fmt.Println(math.Ceil(1.3))

	fmt.Println(math.Max(4, 8))
	fmt.Println(math.Min(4, 8))

}
