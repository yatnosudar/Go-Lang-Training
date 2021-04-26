package main

import (
	"fmt"
)

func main() {
	var nama [3]string
	nama[0] = "Eko"
	nama[1] = "Kurniawan"
	nama[2] = "Khannedy"

	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])

	var angka = [3]int{
		98,
		87,
		76,
	}

	fmt.Println(angka)

	fmt.Println(len(nama))
	fmt.Println(len(angka))
	angka[0] = 100
	fmt.Println(angka)

	var tes [8]string
	fmt.Println(len(tes))

}
