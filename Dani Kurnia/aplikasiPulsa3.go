package main

import "fmt"

type Pemesan struct {
	NoHp    string
	Nominal map[int]int
}

func main() {

	var customer1 Pemesan
	customer1.order(10000, 25000, 15000)
}

func (pemesan Pemesan) order(nominal ...int) {
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
	pemesan.NoHp = "0821883122"
	fmt.Println("\nAnda memesan dengan No Hp :", pemesan.NoHp)
	fmt.Println("Dengan pengisian saldo", nominal[0])
	fmt.Println("Dengan harga :", pemesan.Nominal[10000])

	pemesan.NoHp = "0882163921"
	fmt.Println("\nAnda memesan dengan No Hp :", pemesan.NoHp)
	fmt.Println("Dengan pengisian saldo", nominal[1])
	fmt.Println("Dengan harga :", pemesan.Nominal[25000])

	pemesan.NoHp = "+62116524372"
	fmt.Println("\nAnda memesan dengan No Hp :", pemesan.NoHp)
	fmt.Println("Dengan pengisian saldo", nominal[2])
	fmt.Println("Dengan harga :", pemesan.Nominal[15000])
}
