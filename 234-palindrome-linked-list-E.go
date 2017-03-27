//https://leetcode.com/problems/palindrome-linked-list/?tab=Description

// 1 find mid node,2 reverse high half list,3 compare

func isPalindrome(head *ListNode) bool {
    if head==nil || head.Next==nil {
        return true
    }
    var fast,slow=head,head
    var p,r*ListNode
    for fast.Next!=nil&&fast.Next.Next!=nil{
        fast,slow=fast.Next.Next,slow.Next
    }
    p,slow.Next=slow.Next,nil
    r=nil
    for p!=nil{
        q:=p.Next
        p.Next,r=r,p
        p=q
    }
    p=head
    for r!=nil &&r.Val==p.Val{
        r,p=r.Next,p.Next
    }
    return r==nil
 } 
