package main

import (
    "fmt"
)

func removeElement(nums []int, val int) int {
    result := nums[:0]
    for _, v := range nums {
        if v != val {
            result = append(result, v)
        }
    }
    return len(result)
}

func main() {
    fmt.Println(removeElement([]int{3,2,2,3}, 3))
    fmt.Println(removeElement([]int{0,1,2,2,3,0,4,2}, 2))
}
