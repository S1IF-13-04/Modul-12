package main

import "fmt"

func main() {
	var username, password string
	salah := 0
	for username != "Admin" || password != "Admin" {
		fmt.Scan(&username, &password)
		if username != "Admin" || password != "Admin" {
			salah++
		}
	}
	fmt.Println(salah, "percobaan gagal login ")

}
