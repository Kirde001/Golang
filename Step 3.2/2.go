package main

import "fmt"
func main2() {
    var n1 int
    fmt.Scan(&n1)
	if n1 > -1 && n1 < 17 {
		fmt.Println("Принадлежит")
    } else {
		fmt.Println("Не принадлежит")
	}
}