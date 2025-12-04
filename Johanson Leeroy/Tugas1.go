package main

import "fmt"

func main() {
	var username, password string
	var gagal int
	gagal = 0
	fmt.Scan(&username, &password)
	for username != "Admin" || password != "Admin" {
		gagal++
		fmt.Scan(&username, &password)
	}
	fmt.Println(gagal, "Percobaan gagal")
}
