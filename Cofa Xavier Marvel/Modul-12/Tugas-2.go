package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num, mod10, digit int
	fmt.Scan(&num)

	str := strconv.Itoa(num)
	len := len(str)

	mod10 = 10

	for len >= 1 {
		digit = (num % mod10) / (mod10 / 10)
		len--
		mod10 *= 10
		fmt.Print(digit, " ")
	}

}
