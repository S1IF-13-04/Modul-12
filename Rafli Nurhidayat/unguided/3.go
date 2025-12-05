package main

import "fmt"

func main() {
	var x, y, counter int

	fmt.Scan(&x, &y)

	for x >= y {
		x = x - y
		counter++
	}
	fmt.Print(counter)
}