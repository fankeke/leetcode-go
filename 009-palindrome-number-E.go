//https://leetcode.com/problems/palindrome-number/?tab=Description


func isPalindrome(x int) bool {
    x1,x2:=x,0
    for x>0 {
        x2=x2*10+x%10
        x=x/10
    }
    return x1==x2
}
