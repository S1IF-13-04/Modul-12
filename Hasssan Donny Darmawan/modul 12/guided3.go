package main
import "fmt"

func main (){
	var n,a,b,c int
	fmt.Scan(&n)
	i:= 1
	a= 0
	b= 1

	for i<=n && n>=2 {
		fmt.Print(a, " ")
		c= a+b
		a=b
		b=c
		i++
	}
}