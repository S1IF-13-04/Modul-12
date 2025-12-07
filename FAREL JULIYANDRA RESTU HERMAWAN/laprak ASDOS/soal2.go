package main

import "fmt"

func main() {
	var n uint64
	fmt.Scan(&n)
	if n == 0 {
		fmt.Println(0)
		return
	}
	for n > 0 {
		fmt.Println(n % 10)
		n /= 10
	}
}
