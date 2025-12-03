package main

import "fmt"

func main() {
	var passcode, username string
	var failedAttempts int
	fmt.Scan(&passcode, &username)

	for passcode != "admin" || username != "admin" {
		fmt.Scan(&passcode, &username)
		failedAttempts++
	}
	fmt.Println(failedAttempts, "failed login attempts")
}
