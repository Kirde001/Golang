package main

import "fmt"

func main7() {
	var n1 int
	fmt.Scan(&n1)

	if (n1 % 4 == 0 && n1 % 100 != 0) || n1 % 400 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
