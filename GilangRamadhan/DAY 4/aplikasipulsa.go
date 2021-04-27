package main

import "fmt"

func main() {

	fmt.Println("aplikasi pulsa")
	number := "82126525776"
	fmt.Println("NO HP =", number)
	harga := ""
	fmt.Println("isi pulsa =", harga)

	switch number {
	case "082126525776":
		fmt.Println("number sudah benar")
	case "6282126525776":
		fmt.Println("number sudah benar")
	case "82126525776":
		fmt.Println("number yang anda masukan salah")

		if harga == "10.000" {
			fmt.Println("Total Harga RP 12.000")
		}
		if harga == "15.000" {
			fmt.Println("Total Harga RP 17.500")
		}
		if harga == "20.000" {
			fmt.Println("Total Harga RP 21.000")
		}
		if harga == "25.000" {
			fmt.Println("Total Harga RP 27.000")
		}
		if harga == "50.000" {
			fmt.Println("Total Harga RP 55.000")
		}
		if harga == "100.000" {
			fmt.Println("Nominal yang anda masukan tidak terdaftar")
		}

	}

}
