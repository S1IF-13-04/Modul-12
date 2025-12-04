package main

import "fmt"

func main() {
	var bil int
	fmt.Scan(&bil)
	for bil > 1 {
		fmt.Print(bil, " X ")
		bil--
	}
	fmt.Print("1")
}
