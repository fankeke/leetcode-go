//https://leetcode.com/problems/plus-one/?tab=Description

func plusOne(digits []int) []int {
    carry:=1
    for i:=len(digits)-1;i>=0;i--{
        temp:=digits[i]+carry
        digits[i],carry=temp%10,temp/10
    }
    if carry==1{
        digits=append([]int{1},digits...)
    }
    return digits
}
