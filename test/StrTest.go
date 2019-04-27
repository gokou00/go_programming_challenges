package main

import "fmt"

func main() {
	test := "abbz"
	//letter := 0
	final := ""

	for _, x := range test {
	    
	    if (x == 90){
	        final += string(65)
	        continue
	    }
	    
	    if(x == 122){
	        final += string(97)
	        continue
	    }
	    
	    final += string(x + 1)
	    
	    
		//fmt.Println(string(x +1))
	}
	
	fmt.Println(final)

}
