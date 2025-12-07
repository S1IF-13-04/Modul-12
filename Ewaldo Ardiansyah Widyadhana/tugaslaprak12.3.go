package main

import "fmt"

func main() {
	var x int
	var y int
	var result int = 0

	fmt.Print("Masukkan bilangan X (pembilang/dividend): ")
	fmt.Scan(&x)

	fmt.Print("Masukkan bilangan Y (pembagi/divisor, Y harus <= X): ")
	fmt.Scan(&y)

	if x <= 0 || y <= 0 {
		fmt.Println("Error: X dan Y harus bilangan bulat positif.")
		return
	}
	if y > x {
		fmt.Println("Error: Y tidak boleh lebih besar dari X (ketentuan soal X >= Y).")
		return
	}
	for x >= y {
		x = x - y
		result++
	}

	fmt.Printf("\nKeluaran:\n")
	fmt.Printf("Hasil integer division adalah: %d\n", result)
}
