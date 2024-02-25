package main

import (
    "fmt"
)

func convertToTitle(columnNumber int) string {
    columnTitle := ""
    for columnNumber > 0 {
        columnNumber--
        remainder := columnNumber % 26
        columnTitle = string(rune('A'+remainder)) + columnTitle
        columnNumber /= 26
    }

    return columnTitle
}

func main() {
    fmt.Println(convertToTitle(1))
    fmt.Println(convertToTitle(28))
    fmt.Println(convertToTitle(701))
}
