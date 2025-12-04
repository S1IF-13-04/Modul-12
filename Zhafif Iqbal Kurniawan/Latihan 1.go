package main

import "fmt"

func main() {
	var user, pw string
	fmt.Scan(&user, &pw)
	i := 0
	for user != "Admin" || pw != "Admin" {
		i++
		fmt.Scan(&user, &pw)
	}
	fmt.Println(i, "percobaan gagal login")
}
