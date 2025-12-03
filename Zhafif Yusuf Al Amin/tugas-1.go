package main

import "fmt"

func main() {
	var u, p string
	fmt.Scan(&u, &p)
	ux := "Admin"
	px := "Admin"
	gagal := 0
	for u != ux || p != px {
		gagal++
		fmt.Scan(&u,&p)
	}
	fmt.Println(gagal)
	fmt.Print("percobaan gagal login ")
}
