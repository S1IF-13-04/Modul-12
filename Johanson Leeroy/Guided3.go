package main

import "fmt"

func main() {
	var bil, x, y, i, hasil int
	fmt.Scan(&bil)
	x = 0
	y = 1
	i = 3
	fmt.Print(x, " ", y, " ")
	for i <= bil {
		hasil = x + y
		fmt.Print(hasil, " ")
		x = y
		y = hasil
		i++
	}
}
