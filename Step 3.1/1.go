package main

import "fmt"
func main1() {
    var n1,n2 int
    fmt.Scan(&n1)
    fmt.Scan(&n2)
    if n1 > n2 {
        fmt.Println(n1)
    } else {
        fmt.Println(n2)
    }
}
