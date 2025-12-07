package main

import (
	"fmt"
)

func main() {
	var user, pass string
	gagal := 0

	for {
		fmt.Scan(&user, &pass)

		if user == "Admin" && pass == "Admin" {
			fmt.Println(gagal, "percobaan gagal login")
			break
		} else {
			gagal++
		}
	}
}
