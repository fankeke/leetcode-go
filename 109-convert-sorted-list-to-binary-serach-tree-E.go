//https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/?tab=Description
func sortedListToBST(head *ListNode) *TreeNode {
    if head==nil {
        return nil
    }else if head.Next==nil{
        return &TreeNode{head.Val,nil,nil}
    }
    fast,slow:=head,head
    var pres *ListNode
    for fast.Next!=nil &&fast.Next.Next!=nil {
        fast=fast.Next.Next
        pres=slow
        slow=slow.Next
    }
    if pres==nil{
        pres=head
        slow=slow.Next
    }
    pres.Next=nil
    p:=slow.Next
    slow.Next=nil
    root:=&TreeNode{slow.Val,nil,nil}
    root.Left=sortedListToBST(head)
    root.Right=sortedListToBST(p)
    return root
}
