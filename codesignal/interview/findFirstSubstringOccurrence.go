func findFirstSubstringOccurrence(s string, x string) int {
    firstLetter := string(x[0])
    
    for i,letter := range s{
        if len(x) > len(s[i:]){
            return -1
        }
        if string(letter) == firstLetter{
            length := len(x)
            if i+length <= len(s){
                //fmt.Println("test")
                if s[i:i+length] == x{
                    return i
                }
            }
        }
    }
    
    return -1
    

}
