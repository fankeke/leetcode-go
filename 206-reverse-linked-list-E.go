//https://leetcode.com/problems/reverse-linked-list/?tab=Description
func reverseList(head *ListNode) *ListNode {
    if head==nil || head.Next==nil {
        return head
    }
    var r *ListNode=nil
    p:=head
    for p!=nil {
        q:=p.Next
        p.Next=r
        r=p
        p=q
    }
    return r
}
