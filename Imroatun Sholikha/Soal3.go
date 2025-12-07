package main

import "fmt"

func main() {
	var x, y, i int
	fmt.Scanln(&x, &y)

	i = 0

	for x >= y {
		x = x - y
		i = i + 1
	}
	fmt.Println(i)
}
