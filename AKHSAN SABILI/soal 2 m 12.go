package main
import "fmt"
func main(){
	var n int
	fmt.Scan(&n)
	for {
		cacah := n % 10
		fmt.Println(cacah," ")
		n /= 10
		if n == 0{
			break
		}
	}
	
}