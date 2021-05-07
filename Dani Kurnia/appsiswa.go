package main

import "fmt"

type Siswa struct {
	Nama, Kelas string
	Nilai       int
}

func main() {
	data(
		Siswa{
			Nama:  "Asep",
			Kelas: "XI RPL 1",
			Nilai: 80,
		},
		Siswa{
			Nama:  "Darma",
			Kelas: "XI RPL 1",
			Nilai: 30,
		},
		Siswa{
			Nama:  "Erlan",
			Kelas: "XI RPL 1",
			Nilai: 10,
		},
		Siswa{
			Nama:  "Gina",
			Kelas: "XI RPL 2",
			Nilai: 50,
		},
		Siswa{
			Nama:  "Indah",
			Kelas: "XI RPL 2",
			Nilai: 90,
		},
		Siswa{
			Nama:  "Jajang",
			Kelas: "XI RPL 2",
			Nilai: 20,
		},
		Siswa{
			Nama:  "Kaka",
			Kelas: "XI RPL 3",
			Nilai: 40,
		},
		Siswa{
			Nama:  "Lala",
			Kelas: "XI RPL 3",
			Nilai: 70,
		},
		Siswa{
			Nama:  "Wira",
			Kelas: "XI RPL 3",
			Nilai: 60,
		},
		Siswa{
			Nama:  "Yaya",
			Kelas: "XI RPL 3",
			Nilai: 100,
		},
	)
}

func data(siswa ...Siswa) {
	fmt.Println("Daftar nilai siswa kelas XI :")

	for _, value := range siswa {

		DataSiswa := value
		Grade := DataSiswa.Nilai

		fmt.Println("\nNama :", DataSiswa.Nama)
		fmt.Println("Kelas :", DataSiswa.Kelas)
		fmt.Println("Nilai :", DataSiswa.Nilai)

		if Grade >= 0 && Grade <= 30 {
			Grade := "D"
			fmt.Println("Grade :", Grade)
		} else if Grade >= 31 && Grade <= 70 {
			Grade := "C"
			fmt.Println("Grade :", Grade)
		} else if Grade >= 71 && Grade <= 80 {
			Grade := "B"
			fmt.Println("Grade :", Grade)
		} else if Grade >= 81 && Grade <= 100 {
			Grade := "A"
			fmt.Println("Grade :", Grade)
		}
	}
}
