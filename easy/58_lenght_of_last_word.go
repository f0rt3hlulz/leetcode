package main

import (
    "fmt"
    "strings"
)

func lengthOfLastWord(s string) int {
    s = strings.TrimSpace(s)
    lenght := 0
    for i := len(s) -1; i >= 0; i-- {
        if s[i] == ' ' {
            break
        }
        lenght++
    }
    return lenght
}

func main() {
    fmt.Println(lengthOfLastWord("Hello World"))
    fmt.Println(lengthOfLastWord("   fly me    to   moon  "))
    fmt.Println(lengthOfLastWord("luffy is still joyboy"))
}
