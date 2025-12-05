package main
import "fmt"
func main() {
	var fn, a,b, c int
	fmt.Scan(&fn)
	a = 0 
	b = 1
	i := 0
	for i < fn{
		fmt.Print(a," ")
		c = a+b
		a = b
		b = c
		i++
	}
}	
