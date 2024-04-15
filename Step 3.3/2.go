package main

import "fmt"

func main2() {
	var n1 int
	fmt.Scan(&n1)
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

