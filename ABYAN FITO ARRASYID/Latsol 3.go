package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	hasil := 0
	total := x
	for total >= y {
		total -= y
		hasil++
	}

	fmt.Println(hasil)
}
