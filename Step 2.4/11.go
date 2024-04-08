package main

import (
    "fmt"
)

func main() {
	var n1 int
	fmt.Scan(&n1)
	fmt.Println("Следующее за числом",n1,"число:",n1+1)
	fmt.Println("Для числа",n1,"предыдущее число:",n1-1)
}