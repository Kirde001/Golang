package main

import (
	"fmt" 
	"math"
)

func main8() {
	var n1,n2,n3,n4 float64
    fmt.Scan(&n1, &n2, &n3, &n4)
	var n5 = math.Sqrt(math.Pow(n1-n3, 2) + math.Pow(n2-n4, 2))
	fmt.Println(n5)
}