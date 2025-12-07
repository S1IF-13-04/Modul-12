package main
import "fmt"

func main() {
	var username, password string
	var salah int
	fmt.Scan(&username, &password)
	for username != "admin" || password != "admin"{
		salah += 1
		fmt.Print("Username dan password salah, ulangi: ")
		fmt.Scan(&username, &password)

	}

	fmt.Println(salah)
	fmt.Println("Login Berhasil")
}