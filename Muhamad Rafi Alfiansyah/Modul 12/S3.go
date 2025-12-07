package main
import "fmt"
func main(){
	var x,y,h int
	fmt.Scan(&x,&y)
	for x >= y{
		h = h + 1
		x = x-y
	}
	fmt.Println(h)
}