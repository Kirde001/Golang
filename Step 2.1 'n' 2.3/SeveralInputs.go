package main

import ("fmt"
        "bufio"
        "os"
)

func main5() {
    var s1,s2,s3 string
    a := bufio.NewScanner(os.Stdin)
    _ = a.Scan()
    s1 = a.Text() 
    
    _ = a.Scan()
    s2 = a.Text() 
    
    _ = a.Scan()
    s3 = a.Text() 
    
    fmt.Println(s1)
    fmt.Println(s2)
    fmt.Println(s3) 
}