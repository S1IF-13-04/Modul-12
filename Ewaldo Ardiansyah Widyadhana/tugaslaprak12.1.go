package main

import "fmt"

func main() {
	var usernameInput string
	var passwordInput string
	var failureCount int = 0

	const correctUsername string = "Admin"
	const correctPassword string = "Admin"

	for {
		fmt.Print("Masukkan Username: ")
		fmt.Scan(&usernameInput)

		fmt.Print("Masukkan Password: ")
		fmt.Scan(&passwordInput)

		if usernameInput == correctUsername && passwordInput == correctPassword {
			fmt.Printf("\nLogin Berhasil!\n")
			fmt.Printf("%d percobaan gagal login\n", failureCount)
			break
		} else {
			failureCount++
			fmt.Println("Username atau Password salah. Silakan coba lagi.\n")
		}
	}
}
