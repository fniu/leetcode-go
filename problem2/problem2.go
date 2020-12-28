package problem2

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l1p, l2p *ListNode
	l1p = l1
	l2p = l2
	l3 := &ListNode{}
	var l3p = l3
	carryOver := 0
	for {
		if l1p == nil && l2p == nil {
			return l3
		} else if l1p == nil {
			l3p.Val += l2p.Val
		} else if l2p == nil {
			l3p.Val += l1p.Val
		} else {
			l3p.Val += l1p.Val + l2p.Val
		}

		carryOver = l3p.Val / 10
		l3p.Val = l3p.Val % 10

		if l1p != nil {
			l1p = l1p.Next
		}
		if l2p != nil {
			l2p = l2p.Next
		}
		if l1p != nil || l2p != nil || carryOver != 0 {
			l3p.Next = &ListNode{Val: carryOver}
			l3p = l3p.Next
		}

	}
}

func AddTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	var l1p, l2p *ListNode
	l1p = l1
	l2p = l2
	carryOver := 0
	for {
		if l1p == nil {
			return l1
		}
		if l2p != nil {
			l1p.Val += l2p.Val
		}

		carryOver = l1p.Val / 10
		l1p.Val = l1p.Val % 10

		if l1p.Next == nil && l2p != nil && l2p.Next != nil {
			l1p.Next = l2p.Next
			l1p = l1p.Next
			l1p.Val += carryOver
			l2p = nil
		} else {
			if l2p != nil {
				l2p = l2p.Next
			}
			if l1p.Next == nil && carryOver > 0 {
				l1p.Next = &ListNode{Val: carryOver}
			} else if l1p.Next != nil {
				l1p.Next.Val += carryOver
			}
			l1p = l1p.Next
		}
	}
}
