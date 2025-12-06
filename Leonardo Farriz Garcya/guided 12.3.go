package main

import "fmt"

func main() {
    var N int
    fmt.Scan(&N)

    a := 0
    b := 1

    for i := 1; i <= N; i++ {
        fmt.Print(a, " ")
        c := a + b
        a = b
        b = c
    }
}
