package main

import (
	"fmt"
	"strings"
)

func pesan() {
	hargaPulsa := map[int]int{
		10000: 12000,
		15000: 17500,
		20000: 21000,
		25000: 21000,
		50000: 55000,
	}
	fmt.Println("")
	fmt.Println("\nDaftar harga pulsa : ")
	for key, value := range hargaPulsa {
		fmt.Println(key, "=", value)
	}
}
func MasukanNo(number ...string) {
	var NumberYangMasuk = strings.Join(number, ", ")
	fmt.Println("Number anda sudah benar: ", NumberYangMasuk)
}

func MasukanNominal(nominal ...string) {
	var JumlahNominal = strings.Join(nominal, ", ")
	fmt.Println("Nominal yang anda masukan sudah benar : ", JumlahNominal)
}
func main() {
	pesan()
	MasukanNo("082126525776", "08122411776", "08123346972")
	MasukanNominal("10000 Total Harga = RP 12.000", "20000 Total Harga = RP 22.000", "50000 Total Harga = RP 55.000")

}
