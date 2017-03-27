//https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/?tab=Description
func deleteDuplicates(head *ListNode) *ListNode {
    if head==nil {
        return head
    }
    p,q:=head,head.Next
    for q!=nil {
        if p.Val==q.Val{
            p.Next,q=q.Next,q.Next
        }else{
            p,q=q,q.Next
        }
    }
    return head
}
