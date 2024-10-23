package main

import "fmt"

func main() {
	var F, C int
	fmt.Print("Masukan fahrenheit:")
	fmt.Scan(&F)
	C = (F - 32) * 5 / 9
	fmt.Println("Hasilnya adalah ", C)
}
