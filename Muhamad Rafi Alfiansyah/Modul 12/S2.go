package main
import "fmt"
func main(){
	var n int
	fmt.Scan(&n)
	for n > 0{
		s := n % 10
		n = n / 10
		fmt.Println(s)
	}
}