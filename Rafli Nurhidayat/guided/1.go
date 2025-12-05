package main

import "fmt"

func main() {
	var x int

	fmt.Scan(&x)

	if x >= 0 {
		for x > 1 {
			fmt.Print(x, " x ")
			x--
		}
		fmt.Print("1")
	}
}
