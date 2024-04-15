package main

import "fmt"
func main() {
    var n1 int
    fmt.Scan(&n1)
    if n1 > 0 {
        fmt.Println(1)
    } else {
        if n1 == 0 {
			fmt.Println(0)
		}else {
			fmt.Println(-1)
    	}
	}
}