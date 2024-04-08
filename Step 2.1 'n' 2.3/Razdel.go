package main

import (
    "fmt"
    "os"
    "bufio"
)

func main7() {
    var v1,v2,v3,v4 string
    s := bufio.NewScanner(os.Stdin)
    
    _ = s.Scan()
    v1 = s.Text()
    
    _ = s.Scan()
    v2 = s.Text()
    
    _ = s.Scan()
    v3 = s.Text()
    
    _ = s.Scan()
    v4 = s.Text()
    
    fmt.Println(v2+v1+v3+v1+v4)
}