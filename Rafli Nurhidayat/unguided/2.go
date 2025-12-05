package main

import "fmt"

func main() {
	var x int

	fmt.Scan(&x)

	sisa := x
	for sisa > 0 {
		terakhir := sisa % 10
		fmt.Println(terakhir)
		sisa = sisa / 10
	}
}