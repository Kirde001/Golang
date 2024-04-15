package main

import "fmt"
func main4() {
    var n1 int
    fmt.Scan(&n1)
    if n1 % 2 == 0 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}