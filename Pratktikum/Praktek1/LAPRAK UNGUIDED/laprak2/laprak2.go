package main

import "fmt"

func main() {
	var (
		nama, nim, kelas string
		prodi            = "S1-12F"
	)
	fmt.Print("Masukan Nama: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukan NIM: ")
	fmt.Scanln(&nim)

	fmt.Print("Masukan Kelas: ")
	fmt.Scanln(&kelas)

	fmt.Println("Perkenalkan saya adalah", nama, "salah satu mahasiswa dari prodi", prodi, "dari kelas", kelas, "dengan NIM", nim)
}
