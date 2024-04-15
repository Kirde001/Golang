package main

import "fmt"
func main5() {
    var n1,n2 int
    fmt.Scan(&n1)
	fmt.Scan(&n2)
    if n1 % n2 == 0 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}