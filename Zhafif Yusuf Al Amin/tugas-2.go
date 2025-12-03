package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	h:=0
	for x > 0 {
		h = x % 10
		fmt.Println(h)
		x /= 10

	}
}
