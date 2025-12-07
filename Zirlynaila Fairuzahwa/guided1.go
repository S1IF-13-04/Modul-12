package main
import "fmt"

func main(){
	var bilangan int
	fmt.Scan(&bilangan)
	for bilangan > 1 {
		fmt.Print(bilangan, " x ")
		bilangan--
	}
	fmt.Print(1)
}