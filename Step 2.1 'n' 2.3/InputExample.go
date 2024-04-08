package main

import (
   "bufio"
   "os"
   "fmt"
)

func main2() { // коммент для избежания ошибок с последующими программами
   scanner := bufio.NewScanner(os.Stdin) 
   _ = scanner.Scan() 
   name := scanner.Text() 
   fmt.Println(name) 
}