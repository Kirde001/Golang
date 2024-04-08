package main

import "fmt"

func main2() {
	var n1 int
    fmt.Scan(&n1)
	fmt.Println((n1%100-n1%10)/10)
}