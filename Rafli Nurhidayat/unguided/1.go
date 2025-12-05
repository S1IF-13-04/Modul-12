package main

import "fmt"

func main() {
	var username, password string
	var counter int

	fmt.Scan(&username, &password)

	counter = 0

	for username != "Admin" || password != "Admin" {
		counter += 1
		fmt.Print("Username atau password salah, ulangi: ")
		fmt.Scan(&username, &password)
	}
	fmt.Print(counter, " percobaan gagal login")
}
