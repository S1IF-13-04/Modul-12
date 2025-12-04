package main

import "fmt"

func main() {
	var user, pass string
	salah := 0

	for {
		fmt.Scan(&user, &pass)
		if user == "Admin" && pass == "Admin" {
			break
		} else {
			salah++

		}
	}

	fmt.Println(salah, "percobaan gagal login")
}
