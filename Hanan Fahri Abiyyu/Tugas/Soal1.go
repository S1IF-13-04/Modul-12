package main

import "fmt"

func main() {
	var username, password string
	fmt.Scan(&username, &password)
	i := 0
	for username != "Admin" || password != "Admin" {
		i++
		fmt.Scan(&username, &password)
	}
	fmt.Println(i, "Percobaan gagal login")
}
