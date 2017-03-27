//https://leetcode.com/problems/remove-nth-node-from-end-of-list/?tab=Description
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    p,q:=head,head
    for i:=0;i<n;i++ {
        p=p.Next
    }
    if p==nil{
        return q.Next
    }
    for p.Next!=nil {
        p,q=p.Next,q.Next
    }
    q.Next=q.Next.Next
    
    return head
}
