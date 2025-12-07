package main
import "fmt"
func main() {
	var angka, bagi, hasil int
	fmt.Scan(&angka, &bagi)
	for angka >= bagi {
		angka = angka - bagi
		hasil++
	}
	fmt.Print(hasil)
}