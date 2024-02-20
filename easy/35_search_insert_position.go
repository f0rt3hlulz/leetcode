package main

import (
    "fmt"
)


func searchInsert(nums []int, target int) int {
    if len(nums) == 0 {
        return 0
    }
    for index, value := range nums {
        if target <= value {
            return index
        }
    }
    return len(nums)
}

func main() {
    fmt.Println(searchInsert([]int{1,3,5,6}, 5));
    fmt.Println(searchInsert([]int{1,3,5,6}, 2));
    fmt.Println(searchInsert([]int{1,3,5,6}, 7));
}
