package main

import "fmt"

func pesan() (noHp string, nominal int) {
	noHp = "0881572312"
	nominal = 15000
	hargaPulsa := map[int]int{
		10000: 12000,
		15000: 17500,
		20000: 21000,
		25000: 21000,
		50000: 55000,
	}
	fmt.Println("")
	fmt.Println("Daftar harga pulsa : ")
	for key, value := range hargaPulsa {
		fmt.Println(key, "=", value)
	}
	fmt.Println("\nNo hp anda : ", noHp)
	fmt.Println("Isi pulsa : ", nominal)

	if nominal == 10000 {
		fmt.Println("Total harga : 12.000")
	} else if nominal == 15000 {
		fmt.Println("Total harga : 17.500")
	} else if nominal == 20000 {
		fmt.Println("Total harga : 21.000")
	} else if nominal == 25000 {
		fmt.Println("Total harga : 27.000")
	} else if nominal == 50000 {
		fmt.Println("Total harga : 55.000")
	} else {
		fmt.Println("Daftar pulsa yang ingin anda beli tidak ada")
	}
	return
}

func main() {
	pesan()
}
