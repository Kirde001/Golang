package main

import (
    "fmt"
)

func main9() {
	var n1,n2,n3 int
	fmt.Scan(&n1, &n2, &n3)
	var n4 = n1 * 100 + n2 
	var n5 = n4 * n3
	fmt.Print(n5/100,n5%100)
}