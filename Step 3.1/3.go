package main

import "fmt"
func main3() {
    var n1 int
    fmt.Scan(&n1)
    if n1 >= 12{
        fmt.Println("Доступ разрешен")
    } else {
        fmt.Println("Доступ запрещен")
    }
}