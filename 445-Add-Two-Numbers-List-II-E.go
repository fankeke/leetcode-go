//https://leetcode.com/problems/add-two-numbers-ii/?tab=Description

func reverseList(l*ListNode)*ListNode{
     if l==nil || l.Next==nil {
         return l
     }
     var r *ListNode=nil
     for l!=nil{
         p:=l.Next
         l.Next=r
         r=l
         l=p
     }
     return r
 }
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    p,q:=reverseList(l1),reverseList(l2)
    if p==nil {
        return q
    }else if q==nil {
        return p
    }
    
    var carry=0
    var h ListNode
    var r=&h
    
    for p!=nil || q!=nil {
        sum:=0
        if p!=nil {
            sum,p=sum+p.Val,p.Next
        }
        if q!=nil {
            sum,q=sum+q.Val,q.Next
        }
        sum+=carry
        sum,carry=sum%10,sum/10
        
        n:=&ListNode{sum,nil}
        r.Next=n
        r=n
    }
    if carry==1{
        r.Next=&ListNode{carry,nil}
    }
    
    return reverseList(h.Next)
    
}
