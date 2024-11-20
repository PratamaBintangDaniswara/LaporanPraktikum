package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	if (x%2 == 0) && (x < 0){
		fmt.Print("genap negatif")
	}else{
		fmt.Print("bukan")
	}

}