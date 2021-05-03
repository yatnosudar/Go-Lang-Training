package main

import (
	"GilangRamadhan/DAY8/database"
	 "fmt"
)



func main() {
	result := database.GetDatabase()
	fmt.Println(result)

}
