package main

import (
    "fmt"
    "os"
    "bufio"
)

func main6() {
    var n1,n2,n3 string
    s := bufio.NewScanner(os.Stdin)
    _ = s.Scan()
    n1 = s.Text()
    
    _ = s.Scan()
    n2 = s.Text()
    
    _ = s.Scan()
    n3 = s.Text()
    
    fmt.Println(n3)
    fmt.Println(n2)
    fmt.Print(n1) 
}