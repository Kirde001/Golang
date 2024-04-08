package main

import (
	"fmt" 
	"math"
)

func main() {
	var n1,n2,n3,n4 float64
    fmt.Scan(&n1, &n2, &n3, &n4)
	var n5 = math.Abs(n1-n3) + math.Abs(n2-n4)
	fmt.Println(n5)
}