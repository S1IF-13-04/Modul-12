package main

import "fmt"

func main() {
    var user, pw string
    gagal := 0

    fmt.Scan(&user, &pw)

    for user != "Admin" || pw != "Admin" {
        gagal = gagal + 1
        fmt.Scan(&user, &pw)
    }

    fmt.Printf("%d percobaan gagal login\n", gagal)
}
