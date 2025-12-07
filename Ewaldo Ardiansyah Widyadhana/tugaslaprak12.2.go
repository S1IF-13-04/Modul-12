package main

import "fmt"

func main() {
	var number int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&number)

	if number <= 0 {
		fmt.Println("Input harus bilangan bulat positif.")
		return
	}

	fmt.Println("\nKeluaran:")

	for number > 0 {
		digit := number % 10

		fmt.Println(digit)
		number = number / 10
	}
}
