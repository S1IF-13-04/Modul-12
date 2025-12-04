package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	for x > 0 {
		digit := x % 10
		fmt.Println(digit)
		x = x / 10
	}
}
