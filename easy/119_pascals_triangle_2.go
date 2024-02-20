package main

import "fmt"

func getRow(rowIndex int) []int {
    row := make([]int, rowIndex+1)
    row[0] = 1

    for i := 1; i <= rowIndex; i++ {
        row[i] = int(int64(row[i-1]) * int64(rowIndex-i+1) / int64(i))
    }
    return row
}

func main() {
    fmt.Println(getRow(3))
    fmt.Println(getRow(0))
    fmt.Println(getRow(1))
}
