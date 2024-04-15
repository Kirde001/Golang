package main

import "fmt"

func main5() {
	var n1 int
	fmt.Scan(&n1)
	s1 := n1 / 100000
	s2 := (n1 / 10000) % 10
	s3 := (n1 / 1000) % 10
	s4 := (n1 / 100) % 10
	s5 := (n1 / 10) % 10
	s6 := n1 % 10
	if s1+s2+s3 == s4+s5+s6 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
