package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	var carry int
	var head = &ListNode{}

	var ln = head
	for l1 != nil || l2 != nil {
		v := carry
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}

		if v > 9 {
			v = v - 10
			carry = 1
		} else {
			carry = 0
		}

		ln.Next = &ListNode{Val: v}
		ln = ln.Next
	}

	if carry == 1 {
		ln.Next = &ListNode{Val: 1}
	}

	return head.Next
}
