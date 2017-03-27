//https://leetcode.com/problems/add-two-numbers/?tab=Description
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var head,rear *ListNode
    var promotion int
    
    for l1!=nil||l2!=nil {
        sum:=0
        if l1!=nil{
            sum,l1=sum+l1.Val,l1.Next
        }
        if l2!=nil {
            sum,l2=sum+l2.Val,l2.Next
        }
        sum+=promotion
        sum,promotion=sum%10,sum/10
        
        node:=&ListNode{
            sum,
            nil,
        }
        if head==nil{
            head,rear=node,node
        }else{
            rear.Next=node
            rear=node
        }
    }
    if promotion>0 {
        rear.Next=&ListNode{
            promotion,
            nil,
        }
    }
    return head
}


//or 

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var carray int
    var h=ListNode{0,nil}
    r:=&h
    for l1!=nil || l2!=nil {
        sum:=0
        if l1!=nil {
            sum,l1=sum+l1.Val,l1.Next
        }
        if l2!=nil {
            sum,l2=sum+l2.Val,l2.Next
        }
        sum+=carray
        sum,carray=sum%10,sum/10
        r.Next=&ListNode{sum,nil}
        r=r.Next
    }
    if carray==1 {
        r.Next=&ListNode{carray,nil}
    }
    return h.Next
}
