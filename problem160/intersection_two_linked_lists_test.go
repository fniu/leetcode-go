package problem160

import "testing"

func TestSolution(t *testing.T) {
	var listA, listB, listC *ListNode

	listA = LinkedListFromArray([]int{10, 3, 4, 5})
	listB = LinkedListFromArray([]int{1, 11})
	listB.Next.Next = listA.Next
	listC = getIntersectionNode(listA, listB)
	t.Log(listC.Val)

	listA = LinkedListFromArray([]int{10, 3, 4, 5})
	listB = LinkedListFromArray([]int{1, 11})
	listC = getIntersectionNode(listA, listB)
	if listC != nil {
		t.Errorf("Expect null, got=%#v", listC)
	}

}

func TestLinkedListBuilder(t *testing.T) {
	var l *ListNode
	a := []int{10, 3, 4, 5}
	l = LinkedListFromArray(a)
	t.Log(l.Val, l.Next.Val, l.Next.Next.Val, l.Next.Next.Next.Val)
}

func LinkedListFromArray(a []int) *ListNode {
	l := &ListNode{Val: a[0]}
	head := l
	for i := 1; i < len(a); i++ {
		l.Next = &ListNode{Val: a[i]}
		l = l.Next
	}
	return head
}
