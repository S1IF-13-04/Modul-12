package main

import "fmt"

func main() {
	const usernameBenar = "Admin"
	const passwordBenar = "Admin"

	var user, pass string
	gagal := 0

	for {
		fmt.Scan(&user, &pass)

		if user == usernameBenar && pass == passwordBenar {
			break
		} else {
			gagal++
		}
	}

	fmt.Println(gagal, "percobaan gagal login")
}
