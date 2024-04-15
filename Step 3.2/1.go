package main

import "fmt"
func main1() {
    var n1 int
    fmt.Scan(&n1)
	if n1 > 99 && n1 < 1000 {
		fmt.Println("YES")
    } else {
		fmt.Println("NO")
	}
}