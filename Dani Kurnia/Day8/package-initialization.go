package main

import (
	"BelajarGolang/Day8/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
