package main
import "fmt"
func main(){
	var user, pw string
	salah := 0
	for {
		fmt.Scan(&user, &pw)
		if user == "Admin" && pw == "Admin"{
			break
		}
		salah++
	}
	fmt.Println(salah, "percobaan gagal login")
}