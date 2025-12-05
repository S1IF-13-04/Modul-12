package main
import "fmt"
func main()  {
	var password string
	fmt.Scan(&password)
	i := 0
	for password != "12345abcde"{
		fmt.Println("password salah")
		fmt.Scan(&password)
		i++
	}
	fmt.Println("Selamat anda berhasil login")
}
