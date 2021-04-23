package main

import "fmt"

func main() {
	fmt.Println("Menghitung luas persegi panjang")
	var (
		panjang = 12
		lebar   = 7
		luas    = panjang * lebar
	)
	fmt.Println("Panjang = ", panjang)
	fmt.Println("Lebar = ", lebar)
	fmt.Println("Luas = panjang x lebar")
	fmt.Println("Luas = ", luas)

	fmt.Println("\nMenghitung volume kubus")
	var (
		sisi   = 8
		volume = sisi * sisi * sisi
	)
	fmt.Println("Sisi = ", sisi)
	fmt.Println("Volume = sisi x sisi x sisi")
	fmt.Println("Volume = ", volume)

	fmt.Println("\nMenghitung luas persegi 4 sama sisi")
	var (
		s     = 5
		luasP = s * s
	)
	fmt.Println("Sisi = ", s)
	fmt.Println("Luas = sisi x sisi")
	fmt.Println("Luas = ", luasP)

}
