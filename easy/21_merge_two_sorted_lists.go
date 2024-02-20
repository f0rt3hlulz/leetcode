package main

import (
    "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    tail := dummy

    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            tail.Next = list1
            list1 = list1.Next
        } else {
            tail.Next = list2
            list2 = list2.Next
        }
        tail1 = tail.Next
    }
    if list1 != nil {
        tail.Next = list1
    } else {
        tail.Next = list2
    }

    return dummy.Next
}

func main() {
    fmt.Println(mergeTwoLists([]ListNode{1,2,4}, []ListNode{1,3,4}))
    fmt.Println(mergeTwoLists({}, {})
    fmt.Println(mergeTwoLists({}, {0})
}
