package main
import "fmt"

func main (){
	var pw string
	fmt.Scan(&pw)

	for pw != "12345abcde" {
		fmt.Scan(&pw)
	}
	fmt.Println("Login berhasil")
}