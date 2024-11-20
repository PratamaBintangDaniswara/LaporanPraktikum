package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	if x%2 == 0 {
		motor := x/2
		fmt.Print(motor)
	}else {
		motor := (x/2) + 1
		fmt.Print(motor)
	}
}