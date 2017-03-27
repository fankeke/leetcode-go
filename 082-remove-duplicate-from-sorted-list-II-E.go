//https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/?tab=Description

func deleteDuplicates(head *ListNode) *ListNode {
    if head==nil {
        return head
    }
    var h =&ListNode{0,nil}
    h.Next=head
    prep,p,q:=h,head,head.Next
    for q!=nil{
        if p.Val==q.Val{
            r:=q
            for r!=nil && r.Val==q.Val {
                r=r.Next
            }
            if r==nil{
                prep.Next=r
                return h.Next
            }
            prep.Next=r
            p=prep.Next
            if p==nil{
                return h.Next
            }else{
                q=p.Next
            }
        }else{
            prep,p,q=p,q,q.Next
        }
    }
    return h.Next
}
