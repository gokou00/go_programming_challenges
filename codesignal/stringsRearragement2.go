package main

import (
	"fmt"
	"strings"
)



func main() {

	inputArray := []string{"f", "g","a","h"}
	
	fmt.Println(inputArray)


	if len(inputArray) == 1 {
		fmt.Println("false")
		return
	}
//fmt.Println(inputArray)	

 countequal := 0

	for i := 0; i < len(inputArray)-1; i++ {
		if inputArray[i] == inputArray[i+1] {
			countequal++
		}
	}

	if countequal == len(inputArray)-1 {
		fmt.Println("false")
		return
	}



	
//	fmt.Println("test")
	
    N := len(inputArray) 
    
    p := make([]int,N)
    
    //fmt.Println(p)
    
    for i:= 0; i < len(p) -1; i++{
        p[i] = i
    }
    
    //fmt.Println(p)
   
   
    
    i := 1
    var j int

    for i < N{
        if(p[i] < i){
            if(i % 2 == 1){
                j = p[i]
            }else{
                j = 0
            }
         inputArray[j],inputArray[i] = inputArray[i],inputArray[j]
         //fmt.Println(inputArray)
         
         final := holderf(inputArray)
         
         if(final == true){
             fmt.Println("true")
             return
         }
         
         p[i]++
         i = 1 
        }else{
            p[i] = 0
            i++
        }
        
        
        
 }    
    


fmt.Println("false")

}

func holderf(a []string) bool {

	count := 0
	for i := 0; i < len(a)-1; i++ {
		test := strTest(a[i], a[i+1])
		//fmt.Println(test)

		if test == true {
			count++
		}

	}
	if count == len(a)-1 {
		return true
	}

	return false
}

func strTest(a string, b string) bool {
	count := 0

	for _, x := range a {
		test := strings.ContainsAny(b, string(x))
		if test == false {
			count++
		}
	}

	if count > 1 {
		return false
	}

	count = 0

	for _, x := range b {
		test := strings.ContainsAny(a, string(x))
		if test == false {
			count++
		}
	}

	if count > 1 {
		return false
	}

	//return true
	count = 0

	for i := 0; i < len(a); i++ {
		if string(a[i]) != string(b[i]) {
			count++
		}
	}

	//	fmt.Println(count)

	if count != 1 {
		return false
	}

	return true

}
