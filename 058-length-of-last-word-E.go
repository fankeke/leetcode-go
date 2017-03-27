//https://leetcode.com/problems/length-of-last-word/?tab=Description

func lengthOfLastWord(s string) int {
    if len(s)==0 {
        return 0
    }

    ans:=0
    for i:=len(s)-1;i>=0;i--{
        if s[i]==' '&& ans!=0 {
            return ans
        }
        if s[i]!=' '{
            ans++
        }
    }
}
