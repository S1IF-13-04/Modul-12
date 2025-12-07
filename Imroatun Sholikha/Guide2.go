package main
import "fmt"
func main(){
	var x string
	fmt.Println("Masukan pw :")
	fmt.Scan(&x)

	for x != "12345abcde"{
		fmt.Print("Pasword salah: ")
		fmt.Scan(&x)
	}
	fmt.Println("Selamat anda berhasil")
}