package main

import "fmt"

var hargaPulsa = map[int]int{
	10000: 12000,
	15000: 17500,
	20000: 22000,
	25000: 27000,
	50000: 55000,
}

type Customer struct {
	NoHp    string
	Nominal int
}

//variadic
func main() {
	Pesanan(
		Customer{
			NoHp:    "088212389",
			Nominal: 15000,
		},
		Customer{
			NoHp:    "088212389",
			Nominal: 15000,
		},
		Customer{
			NoHp:    "088212389",
			Nominal: 15000,
		},
		Customer{
			NoHp:    "088212389",
			Nominal: 15000,
		},
		Customer{
			NoHp:    "088212389",
			Nominal: 15000,
		},
		Customer{
			NoHp:    "088212389",
			Nominal: 15000,
		},
		Customer{
			NoHp:    "088212389",
			Nominal: 15000,
		})
}

func Pesanan(order ...Customer) {

	fmt.Println("Daftar harga :")
	for key, value := range hargaPulsa {
		fmt.Println(key, "=>", value)
	}
	for key, value := range order {
		fmt.Println(key, "=>", value)
		pesan := value
		value1, TotalHarga1 := hargaPulsa[pesan.Nominal]
		fmt.Println("\nAnda memesan pulsa dengan No Hp :", pesan.NoHp)
		fmt.Println("Dengan pengisian saldo :", pesan.Nominal)
		if TotalHarga1 {
			fmt.Println("Dengan Harga :", value1)
		}
	}
}

// Guru, mau masukin nilai
// object siswa, Nama, Kelas, Nilai (10,20,30,40,50,60,70,80,90,100)
// Guru, input object siswa
// programnya ngeprint nilai siswa (10~30 => D, 31~70 => C, 71 ~ 80 => B, >81 => A)
