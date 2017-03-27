//https://leetcode.com/problems/sort-list/?tab=Description
 func mergelist(p1,p2 *ListNode)*ListNode{
   if p1==nil {
       return p2
   }else if p2==nil {
       return p1
   }
   var h ListNode
   r:=&h
   for p1!=nil && p2!=nil {
       if p1.Val < p2.Val {
           r.Next,p1=p1,p1.Next
       }else{
           r.Next,p2=p2,p2.Next
       }
       r=r.Next
   }
   if p1!=nil{
       r.Next=p1
   }else{
       r.Next=p2
   }
   return h.Next
 }
func sortList(head *ListNode) *ListNode {
    if head==nil || head.Next==nil {
        return head
    }
    fast,slow:=head,head
    for fast.Next!=nil && fast.Next.Next!=nil {
        fast=fast.Next.Next
        slow=slow.Next
    }
    p:=slow.Next
    slow.Next=nil
    
    p1:=sortList(head)
    p2:=sortList(p)
    
    return mergelist(p1,p2)
}
