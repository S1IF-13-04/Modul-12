package main

import "fmt"

func main() {
	var user, pass string
	kesalahan := 0
	fmt.Println("masukan userName dan password")
	fmt.Scan(&user, &pass)
	for user != "Admin" || pass != "Admin" {
		fmt.Scan(&user, &pass)
		kesalahan += 1
	}
	fmt.Println("login berhasil")
	fmt.Println("percobaan login: ", kesalahan)
}
