package main

import (
    "fmt"
    "strings"
)

func addBinary(a string, b string) string {
   var result strings.Builder
    carry := 0
    i, j := len(a)-1, len(b)-1

    for i >= 0 || j >= 0 || carry > 0 {
        sum := carry
        if i >= 0 {
            sum += int(a[i] - '0')
            i--
        }
        if j >= 0 {
            sum += int(b[j] - '0')
            j--
        }

        carry = sum / 2
        result.WriteByte('0' + byte(sum%2))
    }
    resultStr := result.String()
    return reverseString(resultStr)
}

func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
    fmt.Println(addBinary("11", "1"))
    fmt.Println(addBinary("1010", "1011"))
}
