//import "sort"

func guno(playercount int, moves []string) []int {
    
    var cards map[string]int
    cards = make(map[string]int)
    
    var players map[string]int
    players = make(map[string]int)
    
    var playerArryString []string
    
    playerArryString = make([]string,playercount)
    
    
    cards["C"] = 0
    cards["s1"] = 20
    cards["s2"] = 20
    cards["s3"] = 20
    cards["s4"] = 50
    cards["s5"] = 50
    
   // fmt.Println(cards)
    
    playerNum := " "
    count := 0
    
    for i := 1; i<=playercount;i++{
        playerNum = "p" + strconv.Itoa(i)
        
        playerArryString[count] = playerNum
        count++
       // fmt.Println(playerNum)
        players[playerNum] = 0
    }
    
   // fmt.Println(players)
  //  fmt.Println(playerArryString)
    
    count = 0
    reverse := false
    
   // fmt.Println(players)
    
    for _,x := range moves{
        
        switch x{
            case "c":
              players[playerArryString[count]] += cards[x]
            case "s1":
              players[playerArryString[count]] += cards[x]
            
            if(reverse == false){
              if(count >= len(playerArryString)-1){
                count = 0
              }else{
               count++
              }      
         }
                    if(reverse == true){
                      if(count <= 0){
                       count = len(playerArryString) -1
               
                      }else{
            count--
         }                
        }
        
            

            case "s2":
              players[playerArryString[count]] += cards[x]
            case "s3":
              players[playerArryString[count]] += cards[x]  
            
            if(reverse == true){
                reverse = false
            }else{
                reverse = true
            }
            
            

            case "s4":
              players[playerArryString[count]] += cards[x]
            case "s5":
              players[playerArryString[count]] += cards[x]
            
            
        }
        
        if(reverse == true){
            if(count <= 0){
                count = len(playerArryString) -1
                continue
            }
            count--
        }
        
        
        if(reverse == false){
            if(count >= len(playerArryString)-1){
                count = 0
                continue
            }
            count++
        }
        
        
        fmt.Println(players)
                   
        
    }
    
    //fmt.Println(players)
    
    var answer []int
    
    answer = make([]int,playercount)
    
    for i,p := range playerArryString{
        
        answer[i] = players[p]
        
    }
    
    //fmt.Println(answer)
    
    
    
    return answer
}
//create a map for the cards
//create a string slice for the players
//could also create a linked list of maps for the players key could be player name value is total
//
// nested loops the outer to go through the string inner loop to go through the slice of players and perform a action on the map
// 