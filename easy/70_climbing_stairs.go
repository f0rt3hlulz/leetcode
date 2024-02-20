package main

import "fmt"

func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }

    first, second := 1, 2

    for i := 3; i <= n; i++ {
        current := first + second
        first, second = second, current
    }
    return second
}

func main() {
    fmt.Println(climbStairs(2))
    fmt.Println(climbStairs(3))
}
