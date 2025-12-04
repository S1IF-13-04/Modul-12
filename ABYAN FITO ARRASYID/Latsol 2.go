package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for n > 0 {
		fmt.Println(n % 10)
		n /= 10
	}
}