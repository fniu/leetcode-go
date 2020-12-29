package problem160

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	lenA, lenB := 0, 0
	for pA != nil && pB != nil {
		lenA++
		lenB++
		pA, pB = pA.Next, pB.Next
	}
	for pA != nil {
		lenA++
		pA = pA.Next
	}
	for pB != nil {
		lenB++
		pB = pB.Next
	}
	var diff int
	if lenA > lenB {
		diff = lenA - lenB
	} else {
		diff = lenB - lenA
		headA, headB = headB, headA
	}
	for diff != 0 {
		headA = headA.Next
		diff--
	}
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}

func getIntersectionNode2(pA, pB *ListNode) *ListNode {
	if pA == nil || pB == nil {
		return nil
	}

	d := map[*ListNode]bool{}
	for {
		if pA == nil && pB == nil {
			break
		}
		if pA != nil {
			if val, ok := d[pA]; ok {
				if !val {
					return pA
				}
			} else {
				d[pA] = true
			}
			if pA.Next != nil {
				pA = pA.Next
			} else {
				pA = nil
			}

		}
		if pB != nil {
			if val, ok := d[pB]; ok {
				if val {
					return pB
				}
			} else {
				d[pB] = false
			}
			if pB.Next != nil {
				pB = pB.Next
			} else {
				pB = nil
			}
		}
	}
	return nil
}
