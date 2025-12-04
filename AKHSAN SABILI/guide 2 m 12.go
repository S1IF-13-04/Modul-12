package main
import "fmt"
func main(){
	var n string
	fmt.Scan(&n)
	for n != "12345abcde"{
		fmt.Scan(&n)
	}
	fmt.Println("login berhasil")
}