package main

import (
    "fmt" 
    "os" 
    "bufio"
)

func main4() {
   scanner := bufio.NewScanner(os.Stdin) 
   _ = scanner.Scan() 
   bookName := scanner.Text() 
   fmt.Println(bookName,"- лучшая книга!")
}