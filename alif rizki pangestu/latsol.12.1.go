package main

import (
	"fmt"
)

func main() {
	const correctUsername = "Admin"
	const correctPassword = "Admin"

	var username, password string
	gagal := 0

	// while-loop (for dengan kondisi)
	for {
		fmt.Print("Masukkan username: ")
		fmt.Scan(&username)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&password)

		if username == correctUsername && password == correctPassword {
			break // keluar dari loop jika login berhasil
		}

		fmt.Println("Username atau password salah, coba lagi!\n")
		gagal++
	}

	fmt.Println(gagal, "percobaan gagal login")
}
