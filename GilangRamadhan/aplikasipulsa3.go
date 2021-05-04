package main

import "fmt"

type Pemesan struct {
	NoHp    string
	Nominal map[int]int
}

func (pemesan Pemesan) pesan(nominal ...int) {
	pemesan.Nominal = map[int]int{
		10000: 12000,
		15000: 17500,
		20000: 22000,
		25000: 27000,
		50000: 55000,
	}
	fmt.Println("Daftar harga pulsa :")
	for key, value := range pemesan.Nominal {
		fmt.Println(key, "=>", value)
	}
	pemesan.NoHp = "082126525776"
	fmt.Println("\nAnda memesan  \ndengan No Hp :", pemesan.NoHp)
	fmt.Println("Dengan pengisian nominal :", nominal[0])
	fmt.Println("Dengan harga :", pemesan.Nominal[20000])

	pemesan.NoHp = "08124112044"
	fmt.Println("\nAnda memesan  \ndengan No Hp :", pemesan.NoHp)
	fmt.Println("Dengan pengisian nominal", nominal[1])
	fmt.Println("Dengan harga :", pemesan.Nominal[50000])

	pemesan.NoHp = "+621224525776"
	fmt.Println("\nAnda memesan  \ndengan No Hp :", pemesan.NoHp)
	fmt.Println("Dengan pengisian nominal", nominal[2])
	fmt.Println("Dengan harga :", pemesan.Nominal[10000])
}

func main() {

	var pelanggan Pemesan
	pelanggan.pesan(20000, 50000, 10000)
}
