package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	f1 := 0
	f2 := 1
	m := 0
	for m < n {
		fmt.Print(f1, " ")
		k := f1 + f2
		f1 = f2
		f2 = k
		m++
	}
}