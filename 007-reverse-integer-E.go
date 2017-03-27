//https://leetcode.com/problems/reverse-integer/?tab=Description
func reverse(x int) int {
    var flag,ans int
    if x==0 {
        return x
    }else if x<0 {
        flag,x=-1,-x
    }else{
        flag=1
    }
    ans=0
    for x>0 {
        temp:=x%10
        ans=ans*10+temp
        if ans>2147483647 {
            return 0
        }
        x=x/10
    }
    return flag*ans
}
