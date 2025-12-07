package main
import "fmt"
func main() {
	var i int
	fmt.Scan(&i)
	j := i
	for j > 1 {
		fmt.Print(j," x ")
		j--
	}
	fmt.Println(1)
}