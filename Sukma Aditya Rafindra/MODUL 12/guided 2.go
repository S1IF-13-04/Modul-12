package main

import "fmt"

func main() {
	var x string
	fmt.Print("masukan pw: ")
	fmt.Scan(&x)

	for x != "12345abcde" {
		fmt.Println("pw salah masukan lagi: ")
		fmt.Scan(&x)
	}
	fmt.Print("selamat anda berhasil login")
}
