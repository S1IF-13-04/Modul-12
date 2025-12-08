package main

import "fmt"

func main() {
	var username, password string
	kesalahan := 0
	fmt.Scan(&username, &password)
	for username != "Admin" || password != "Admin" {
		fmt.Scan(&username, &password)
		kesalahan += 1
	}
	fmt.Println(kesalahan, "percobaan gagal login")
}
