package main

import (
	"fmt"
	"leetcode/question203/structure"
)

func removeElements(head *structure.ListNode, val int) *structure.ListNode {
	dummyNode := &structure.ListNode{}
	dummyNode.Next = head
	cur := dummyNode
	for (cur.Next != nil) {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyNode.Next
}

/**
输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
*/
func main() {
	input := []int{1, 2, 6, 3, 4, 5, 6}
	dummyNode := &structure.ListNode{}
	cur := dummyNode
	for _, v := range input {
		cur.Next = &structure.ListNode{
			Val: v,
		}
		cur = cur.Next
	}

	fmt.Println(dummyNode.Next)
	fmt.Println(removeElements(dummyNode.Next, 6))
}
