package main

import "fmt"

type person struct {
	name, kelas string
	Nilai       int
}

func main() {
	var allStudents = []person{
		{name: "Joni",
			Nilai: 70,
			kelas: "XI RPL 2"},
		{name: "Asep",
			Nilai: 80,
			kelas: "XI RPL 2"},
		{name: "Susanto",
			Nilai: 30,
			kelas: "XI RPL 2"},
		{name: "Regi",
			Nilai: 50,
			kelas: "XI RPL 2"},
		{name: "Tatang",
			Nilai: 100,
			kelas: "XI RPL 2"},
	}

	fmt.Println("\nDaftar Nilai Siswa Kelas XI")
	for _, student := range allStudents {

		DataStudent := student
		Grade := DataStudent.Nilai

		fmt.Println("\nNama : ", student.name, "\nNilai", student.Nilai, "\nKelas :", student.kelas)

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
