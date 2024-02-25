package main

import (
    "fmt"
)

func reverseBits(num uint32) uint32 {
    var reversed uint32 = 0
    for i := 0; i < 32; i++ {
        reversed <<= 1
        reversed |= num & 1
        num >>= 1
    }
    return reversed
}

func main() {
}
