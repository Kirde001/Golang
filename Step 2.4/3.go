package main

import "fmt"

func main3() {
    var nn int
    fmt.Scan(&nn)
    nn = nn * nn * nn
    nn = nn * nn
    fmt.Println(nn)
}