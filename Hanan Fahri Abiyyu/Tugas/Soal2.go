package main

import "fmt"

func main() {
	var bilx int
	fmt.Scan(&bilx)

	for bilx > 0{
		Dterakhir := bilx%10
		fmt.Println(Dterakhir, " ")
		bilx /= 10
	} 
}
