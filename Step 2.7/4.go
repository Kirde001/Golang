package main

import (
	"fmt"
	"math"
)

func main() {
	var n1,n2 float64
    fmt.Scan(&n1,&n2)
	fmt.Println(math.Sqrt(math.Pow(n1,2)+math.Pow(n2,2)))
}