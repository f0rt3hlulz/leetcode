package main

import (
    "fmt"
    "unicode"
)

func isPalindrome(s string) bool {
    runes := []rune(s)
    fmt.Println(runes)
    for i, j := 0, len(runes)-1; i < j; {
        fmt.Println(i, j)
        if !unicode.IsLetter(runes[i]) && !unicode.IsDigit(runes[i]) {
            i++
            continue
        }
        if !unicode.IsLetter(runes[j]) && !unicode.IsDigit(runes[j]) {
            j--
            continue
        }
        if unicode.ToLower(runes[i]) != unicode.ToLower(runes[j]) {
            return false
        }
        i++
        j--
    }
    return true
}

func main() {
    fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
    fmt.Println(isPalindrome("race a car"))
    fmt.Println(isPalindrome(" "))
    fmt.Println(isPalindrome("0P"))
}
