package main

import "fmt"

func main(){
    fah := 0.0
    fmt.Println("Enter the degree in fahrenheit")
    fmt.Scanf("%f",&fah)
    
    Celsius := (fah - 32) * 5/9
    
    fmt.Printf("The degree in Celsius is %f \n",Celsius)
}
