package main

import (
	"fmt"
	"strings"
)

func main() {
	daftarHarga()
	harga(20000, 10000, 15000)
	noHp("088123322", "088212332", "088524131")
}

func daftarHarga() {
	daftarHarga := map[int]int{
		10000: 12000,
		15000: 17500,
		20000: 22000,
		25000: 27000,
		50000: 55000,
	}
	fmt.Println("Daftar harga : ")
	for key, value := range daftarHarga {
		fmt.Println(key, "=", value)
	}
}

func harga(nominal ...int) {

	fmt.Println("\nAnda pesan", nominal[0], "dengan harga : 22000")
	fmt.Println("Anda pesan", nominal[1], "dengan harga : 12000")
	fmt.Println("Anda pesan", nominal[2], "Dengan harga : 17500")
}
func noHp(no ...string) {
	noPelanggan := strings.Join(no, ", ")
	fmt.Println("\nNo yang anda masukkan yaitu :", noPelanggan)

}
