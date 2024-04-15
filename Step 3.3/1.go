package main

import "fmt"

func main1() {
	var n1 int
	fmt.Scan(&n1)
	if n1>= -3 && n1 <= 1 {
		fmt.Println("YES")
	} else {
		if n1>=5 && n1 <= 9{
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
