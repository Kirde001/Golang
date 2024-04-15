package main

import "fmt"

func main4() {
	var n1 int
	fmt.Scan(&n1)
	s1 := n1 / 100
	s2 := (n1%100 - n1%10) / 10
	s3 := n1 % 10
	if s1 == s2 || s2 == s3 || s3 == s1 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
