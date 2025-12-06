package main

import "fmt"

func main() {
    var a int
    fmt.Scan(&a)

    i := 1

    if a >= i {
        for a >= i {
            fmt.Print(a)
            if a > 1 {     
                fmt.Print(" x ")
            }
            a--
        }
    } else {
        fmt.Print(i)
    }
}
