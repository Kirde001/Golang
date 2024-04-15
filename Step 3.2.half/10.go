package main

import (
	"fmt"
	"math"
)

func main() {
	var n1,n2,n3,n4 float64
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Scan(&n3)
	fmt.Scan(&n4)
	if math.Abs(n3-n1) == math.Abs(n4-n2){
		fmt.Println("YES")
	} else{
		fmt.Println("NO")
	}
}