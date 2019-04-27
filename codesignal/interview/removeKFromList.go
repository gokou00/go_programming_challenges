// Definition for singly-linked list:
// type ListNode struct {
//   Value interface{}
//   Next *ListNode
// }
//



func removeKFromList(l *ListNode, k int) *ListNode {
   x := l
   var prev = new(ListNode)
   prev = nil
   
   for x != nil{
       x_next := x.Next
       if(x.Value == k) {
               if(prev != nil) {
                  prev.Next = x_next;
               } else {
                  l = x_next;
               }
            } else {
               prev = x;
            }
            x = x_next;
         }
      return l 
   }


