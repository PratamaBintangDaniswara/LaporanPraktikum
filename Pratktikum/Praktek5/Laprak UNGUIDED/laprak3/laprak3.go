package main

import "fmt"

func main() {
	var x, y int
	var hasilx, hasily bool
	fmt.Scan(&x,&y)

	hasilx = y%x == 0
	hasily = x%y == 0

	fmt.Println(hasilx)
	fmt.Println(hasily)
}
