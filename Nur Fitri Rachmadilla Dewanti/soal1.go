package main
import "fmt"
func main() {
    var username, password string
    var salah int
    var isLogin = false

    for isLogin == false {
        fmt.Scan(&username, &password)

        if username != "Admin" || password != "Admin" {
            salah = salah + 1
        } else {
            isLogin = true
        }
    }
    fmt.Println(salah)
    fmt.Println("percobaan gagal login")
}