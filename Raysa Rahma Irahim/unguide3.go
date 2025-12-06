package main

import "fmt"

func main() {
	var x int
	var y int
	var hasil int

	hasil = 0
	fmt.Scan(&x, &y)

	for x >= y {
		x = x - y
		hasil = hasil + 1
	}
	fmt.Println(hasil)
}
