package main

import "fmt"

func main1() {
	x := []bool{false, true, false, true, false, true, false, true}
	y := []bool{false, false, true, true, false, false, true, true}
	z := []bool{false, false, false, false, true, true, true, true}
	var result bool
	for i := range x {
		//result = !(x[i] || !y[i]) || y[i] || !z[i] || (!x[i] && !y[i] && z[i])
		//result = !x[i] || y[i] || !z[i] || (!x[i] && !y[i] && z[i])
		result = !(!(!x[i] && y[i]) && (!y[i]&&z[i]) && !(!x[i] && !y[i] && z[i]))
		fmt.Printf("x: %t, y: %t, z: %t, result: %t\n", x[i], y[i], z[i], result)
	}
}
