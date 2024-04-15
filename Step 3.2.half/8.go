package main

import "fmt"

func main8() {
	var n1 int
	fmt.Scan(&n1)

	if n1 <= 13{
		fmt.Println("детство")
	} else {

		if n1 > 13 && n1 <= 24 {
			fmt.Println("молодость")
			} else {
				if n1 > 24 && n1 <= 59 {
					fmt.Println("зрелость")

                } else {
					fmt.Println("старость")
				}
			}
	}	
}
