package main

func twoSumm(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i]+nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return nil
}

func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
    for i, num := range nums {
        complement := target - num
        if j, ok := numMap[complement]; ok {
            return []int{j, i}
        }
        numMap[num] = i
    }
    return nil
}

func main() {
    twoSum([]int{2, 7, 11, 15}, 9)
}
