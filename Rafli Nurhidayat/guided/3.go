package main

import "fmt"

func main() {
	var n, d1, d2, temp, j int

	fmt.Scan(&n)
	d1 = 0
	d2 = 1
	j = 0

	for j < n {
		fmt.Print(d1, " ")
		temp = d1 + d2
		d1 = d2
		d2 = temp
		j = j + 1

	} 

}