package main
import "fmt"

func main() {
	var a,b int64
	fmt.Scan(&a)
	for a > 0 {
		b = a %10
		fmt.Println(b)
		a /= 10
	}
}