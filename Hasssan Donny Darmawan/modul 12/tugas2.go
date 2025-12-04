package main
import "fmt"

func main (){
	var d, digit int
	fmt.Scan(&d)

	for d>0 {
		digit = d%10
		fmt.Println(digit)
		d/=10
	}
}