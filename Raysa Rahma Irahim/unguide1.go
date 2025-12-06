package main

import "fmt"

func main() {
	var username, password string
	var gagal int

	fmt.Scan(&username, &password)
	for username != "Admin" || password != "Admin" {
		gagal++
		fmt.Scan(&username, &password)
	}
	fmt.Println(gagal, "percobaan gagal login")
}
