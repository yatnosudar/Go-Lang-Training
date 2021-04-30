package main

import (
	"BelajarGolang/Day7/bantuan"
	"fmt"
)

func main() {
	bantuan.SayHello("Eko")
	// bantuan.sayGoodBye("Budi")//error

	fmt.Println(bantuan.Application)
	// fmt.Println(bantuan.version) // error
}
