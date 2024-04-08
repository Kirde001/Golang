package main

import "fmt"

func main5() {
	var n1 float64
    fmt.Scan(&n1)
	fmt.Println(n1-float64(int(n1)))
}