package main

import "fmt"

func main() {
	var N, d1, d2, i, temp int
	fmt.Scan(&N)
	d1 = 0
	d2 = 1
	i = 0
	for i < N {
		fmt.Print(d1, " ")
		temp = d1 + d2
		d1 = d2
		d2 = temp
		i ++
	}
}
