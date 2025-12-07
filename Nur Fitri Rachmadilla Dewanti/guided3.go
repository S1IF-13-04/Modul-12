package main  
import "fmt"  
func main() {  
    var n, s1, s2, i, temp int 
    fmt.Scan(&n) 
    s1 = 0 
    s2 = 1 
    i = 0 
    for i < n { 
        fmt.Print(s1," ") 
        temp = s1 + s2 
        s1 = s2 
        s2 = temp 
        i = i + 1 
    } 
} 