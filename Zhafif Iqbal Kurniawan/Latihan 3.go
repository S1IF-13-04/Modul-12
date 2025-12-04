package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	i := 0

	for x >= y {
		x -= y
		i++
	}
	fmt.Print(i)
}
