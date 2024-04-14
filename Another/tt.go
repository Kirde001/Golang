package main

import (
    "fmt"
)

func main() {
    var n, j int
    fmt.Scan(&n)
    marks := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&marks[i])
    }
    bestI := -1
    bestJ := -1
    fives := 0 
    for i := 0; i < n; i++ {
        CURRENT := 0
        len := 0
        for j := i; j < n; j++ {
            if marks[j] == 2 || marks[j] == 3 {
                break
            }
            if marks[j] == 5 {
                CURRENT++
            }
            len++
        }
        if len > fives || (len == fives && i < bestI) {
            bestI = i
            bestJ = j
            fives = CURRENT
        }
        if CURRENT > 0 && (marks[bestJ+1] == 2 || marks[bestJ+1] == 3) {
            bestI = -1
            bestJ = -1
            fives = 0
        }
    }
    if fives == 0 {
        fmt.Println("-1")
    } else {
        if marks[bestI] == 2 || marks[bestJ] == 2 {
            fmt.Println("-1")
        } else {
            fmt.Println(fives)
        }
    }
}