package main

import "fmt"


func main(){
    
    feet := 0.0
    
    fmt.Println("Please enter the number of feet")
    fmt.Scanf("%f",&feet)
    
    meters := feet * .3048
    
    fmt.Printf("the number of meters are %f \n",meters)
    
}
