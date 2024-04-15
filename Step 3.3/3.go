package main

import "fmt"

func main() {
	var n1,n2,n3 int
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Scan(&n3)
	if n1 < 60 {
		fmt.Println("Легкий вес")
	}	else {
			if n1 <64{
				fmt.Println("Первый полусредний вес")
			} else {
				if n1 < 69{
					fmt.Println("Полусредний вес")
				}
			}
		}
	}

