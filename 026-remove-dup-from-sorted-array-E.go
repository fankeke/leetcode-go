//https://leetcode.com/problems/remove-duplicates-from-sorted-array/?tab=Solutions
func removeDuplicates(nums []int) int {
    length:=len(nums)
    if length<=1 {
        return length
    }
    count:=0
    for i:=1;i<length;i++{
        if nums[i]!=nums[count]{
            count++
            nums[count]=nums[i]
        }
    }
    return count+1
}
