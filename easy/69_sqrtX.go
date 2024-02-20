package main

import (
    "fmt"
)

func mySqrt(x int) int {
    if x < 2 {
        return x
    }

    low, high := 1, x/2
    var mid, sqr int
    for low <= high {
        mid = low + (high - low)/2
        sqr = mid * mid
        if sqr == x {
            return mid
        } else if sqr < x {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return high
}

func main() {
    fmt.Println(mySqrt(4))
    fmt.Println(mySqrt(8))
}
