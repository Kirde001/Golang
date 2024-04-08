package main

import (
    "fmt"
)

func main10() {
	var n1 int
	fmt.Scan(&n1)
	fmt.Println(n1,"мин - это",n1/60,"часов",n1%60,"минут.")
}