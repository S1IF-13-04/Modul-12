package main

import "fmt"

func main() {
	var username, password string
	var failedCount int = 0

	fmt.Scan(&username, &password)
	for username != "Admin" || password != "Admin" {
		failedCount++
		fmt.Scan(&username, &password)
	}
	fmt.Println(failedCount, "percobaan gagal login")
}