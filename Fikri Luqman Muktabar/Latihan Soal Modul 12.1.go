package main

import "fmt"

func main() {
	var username, password string
	i := 0

	fmt.Scan(&username, &password)

	for username != "Admin" || password != "Admin" {
		i++
		fmt.Scan(&username, &password)
	}

	fmt.Println(i, "percobaan gagal login")
}