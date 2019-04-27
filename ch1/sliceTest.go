package main

import "fmt"


func main(){
    x := make([]float64,3,9)
    
    x = append(x,4,5,5,5,5,5,5,5)
    
    fmt.Println(len(x))
    
}