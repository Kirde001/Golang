package main

import "fmt"
func main3() {
    var n1 int
    fmt.Scan(&n1)
	if n1 <= -3 || n1 >= 7 {
		fmt.Println("Принадлежит")
    } else {
		fmt.Println("Не принадлежит")
	}
}