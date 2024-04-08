package main

import (
	"fmt" 
	"math"
)

func main7() {
	var n1,n2 float64
    fmt.Scan(&n1, &n2)
	var n3 = math.Sqrt(math.Pow(n1, 2) + math.Pow(n2, 2))
	fmt.Println(n1+n2+n3)
}