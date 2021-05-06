package main

import "fmt"

type Customer struct {
	NoHp1, NoHp2, NoHp3          string
	Nominal1, Nominal2, Nominal3 int
	Harga                        map[int]int
}

func main() {
	Pesanan(Customer{})
}

func Pesanan(order ...Customer) {
	var pesan = Customer{
		NoHp1:    "088212389",
		NoHp2:    "088767347",
		NoHp3:    "+6231345698",
		Nominal1: 20000,
		Nominal2: 10000,
		Nominal3: 15000,
		Harga: map[int]int{
			10000: 12000,
			15000: 17500,
			20000: 22000,
			25000: 27000,
			50000: 55000,
		},
	}
	fmt.Println("Daftar harga :")
	for key, value := range pesan.Harga {
		fmt.Println(key, "=>", value)
	}

	value1, TotalHarga1 := pesan.Harga[pesan.Nominal1]

	fmt.Println("\nAnda memesan pulsa dengan No Hp :", pesan.NoHp1)
	fmt.Println("Dengan pengisian saldo :", pesan.Nominal1)
	if TotalHarga1 {
		fmt.Println("Dengan Harga :", value1)
	}
	// fmt.Println("Total harga :", pesan.Harga[20000])

	value2, TotalHarga2 := pesan.Harga[pesan.Nominal2]

	fmt.Println("\nAnda memesan pulsa dengan No Hp :", pesan.NoHp2)
	fmt.Println("Dengan pengisian saldo :", pesan.Nominal2)
	if TotalHarga2 {
		fmt.Println("Dengan Harga :", value2)
	}
	// fmt.Println("Total harga :", pesan.Harga[10000])

	value3, TotalHarga3 := pesan.Harga[pesan.Nominal3]

	fmt.Println("\nAnda memesan pulsa dengan No Hp :", pesan.NoHp3)
	fmt.Println("Dengan pengisian saldo :", pesan.Nominal3)
	if TotalHarga3 {
		fmt.Println("Dengan Harga :", value3)
	}
	// fmt.Println("Total harga :", pesan.Harga[15000])
}
