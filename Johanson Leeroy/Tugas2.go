package main

import "fmt"

func main() {
	var bil, digit int
	fmt.Scan(&bil)

	for bil > 0 {
		digit = bil % 10
		bil /= 10
		fmt.Println(digit)
	}
}
