package main

import (
    "fmt"
    "strings"
)

func strStr (haystack string, needle string) int {
    index := strings.Index(haystack, needle)
    return index
}

func main() {
    fmt.Println(strStr("sadbutsad", "sad"))
    fmt.Println(strStr("leetcode", "leeto"))
}
