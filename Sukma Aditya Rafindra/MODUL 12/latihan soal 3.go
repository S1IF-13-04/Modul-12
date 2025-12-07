package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukan x dan y: ")
	fmt.Scan(&x, &y)
	hasil := 0
	temp := x
	for temp >= y {
		temp = temp - y
		hasil++
	}
	fmt.Println(hasil)
}
