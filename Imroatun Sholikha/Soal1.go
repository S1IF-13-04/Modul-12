package main

import "fmt"

func main() {
	var username, password string
	var gagal int
	gagal = 0
	userBenar := "Admin"
	passBenar := "Admin"
	fmt.Scan(&username, &password)

	for username != userBenar || password != passBenar {
		gagal = gagal + 1
		fmt.Scan(&username, &password)
	}
	fmt.Println(gagal, "percobaan gagal login")
}
