package main

import "fmt"

func main() {
	var X, Y, Z int
	fmt.Scan(&X, &Y)

	for X >= Y {
		X = X - Y
		Z++
	}
	fmt.Println(Z)
}
