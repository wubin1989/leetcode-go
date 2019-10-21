package structure

import (
	"bytes"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (self ListNode) String() string {
	var buf bytes.Buffer
	buf.Write([]byte(fmt.Sprint(self.Val)))
	if self.Next == nil {
		return buf.String()
	}
	buf.Write([]byte(", "))
	cur := self
	for (cur.Next != nil) {
		buf.Write([]byte(fmt.Sprint(cur.Next.Val)))
		cur = *cur.Next
		if cur.Next != nil {
			buf.Write([]byte(", "))
		}
	}
	return buf.String()
}
