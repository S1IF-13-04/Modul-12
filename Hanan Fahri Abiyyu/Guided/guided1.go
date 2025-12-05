package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	for n > 1 {
		fmt.Print(n, " ")
		n--
	}
	fmt.Print(1)
}
