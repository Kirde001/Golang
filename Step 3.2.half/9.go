package main

import "fmt"

func main9() {
	var n1,n2,n3,n4 int
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Scan(&n3)
	fmt.Scan(&n4)
	if n1 == n3 || n2 == n4 {
		fmt.Println("YES")
	} else{
		fmt.Println("NO")
	}
}