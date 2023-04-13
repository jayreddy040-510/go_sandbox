package main

func main(nums []int, target int) []int {
}

func twoSum(nums []int, target int) []int {
    ret := make(map[int]int, len(nums) - 1)
    i := 0
    for i < len(nums) {
        ret[nums[i]] = i
        i++
    }
    for i, num := range nums {
        complement := target - num
        if j, ok := ret[complement]; ok && i != j {
            return []int{i, j}
        }
    }
    return []int{-1,-1}
}
