package main

import (
    "fmt"
)

func simpleNumber(nums []int) int {
    result := 0
    for _, num := range nums {
        result ^= num
    }
    return result
}

func main() {
    fmt.Println(simpleNumber([]int {2,2,1}))
    fmt.Println(simpleNumber([]int {4,1,2,1,2}))
}
