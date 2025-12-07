package main
import "fmt"
func main(){
	var user, pw string
	fmt.Scan(&user,&pw)
	gagal := 0
	for user!= "Admin" || pw!="Admin" {
		fmt.Scan(&user,&pw)
		gagal ++
	}
	fmt.Println(gagal,"Percobaan gagal login")
}