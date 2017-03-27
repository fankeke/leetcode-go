//https://leetcode.com/problems/climbing-stairs/?tab=Description
func climbStairs(n int) int {
    temp1,temp2,temp3:=1,2,0
    if n==1 {
        return temp1
    }else if n==2 {
        return temp2
    }else{
        
        for i:=3;i<=n;i++{
            temp3=temp1+temp2
            temp1,temp2=temp2,temp3
        }
        return temp3
    }
}
