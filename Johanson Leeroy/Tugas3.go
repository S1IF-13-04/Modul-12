package main

import "fmt"

func main() {
	var x, y, div int
	fmt.Scan(&x, &y)
	div = 0
	for x >= y {
		x = x - y
		div++
	}
	fmt.Print(div)
}
