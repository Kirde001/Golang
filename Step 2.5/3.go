package main

import "fmt"

func main3() {
	var n1 int
    fmt.Scan(&n1)
	fmt.Println(n1%10+n1/100+(n1%100-n1%10)/10)
}