package main

import "fmt"
func main2() {
    var n1,n2 string
    fmt.Scan(&n1)
    fmt.Scan(&n2)
    if n1 == n2 {
        fmt.Println("Пароль принят")
    } else {
        fmt.Println("Пароль не принят")
    }
}
