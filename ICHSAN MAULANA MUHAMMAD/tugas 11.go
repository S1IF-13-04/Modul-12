package main

import "fmt"

func main() {
	var username, password string
	gagal := 0

	for {
		fmt.Scan(&username, &password)

		if username == "Admin" && password == "Admin" {
			break
		}

		gagal = gagal + 1
	}

	fmt.Println(gagal, "percobaan gagal login")
}
