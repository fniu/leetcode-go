package problem4

import (
	"testing"
)

func TestSolution(t *testing.T) {
	nums1 := []int{2, 2}
	nums2 := []int{2, 3, 10, 10, 11}
	m := findMedianSortedArrays(nums2, nums1)
	if m != 3 {
		t.Errorf("Expect 3, got=%f", m)
	}
	nums1 = []int{2, 2}
	nums2 = []int{2, 3, 10, 10}
	m = findMedianSortedArrays(nums2, nums1)
	if m != 2.5 {
		t.Errorf("Expect 2.5, got=%f", m)
	}
	nums1 = []int{}
	nums2 = []int{2, 3}
	m = findMedianSortedArrays(nums2, nums1)
	if m != 2.5 {
		t.Errorf("Expect 2.5, got=%f", m)
	}
	nums1 = []int{1}
	nums2 = []int{2}
	m = findMedianSortedArrays(nums2, nums1)
	if m != 1.5 {
		t.Errorf("Expect 1.5, got=%f", m)
	}

	nums1 = []int{1}
	nums2 = []int{2, 3}
	m = findMedianSortedArrays(nums2, nums1)
	if m != 2 {
		t.Errorf("Expect 2, got=%f", m)
	}

	nums1 = []int{2}
	nums2 = []int{1, 3, 4}
	m = findMedianSortedArrays(nums2, nums1)
	if m != 2.5 {
		t.Errorf("Expect 2.5, got=%f", m)
	}

	// m := findMedianSortedArrays(nums1, nums2)
	m = findMedianSortedArrays(nums1, nums2)

	nums1 = []int{2, 2, 10}
	nums2 = []int{2, 3, 4}

	// m := findMedianSortedArrays(nums1, nums2)
	m = findMedianSortedArrays(nums1, nums2)

	nums1 = []int{1}
	nums2 = []int{2, 10, 41}
	m = findMedianSortedArrays(nums1, nums2)
	if m != 6 {
		t.Errorf("Expect 6, got=%f", m)
	}

	nums1 = []int{1, 1}
	nums2 = []int{2, 3, 4, 5}

	// m := findMedianSortedArrays(nums1, nums2)
	m = findMedianSortedArrays(nums1, nums2)
	if m != 2.5 {
		t.Errorf("Expect 2.5, got=%f", m)
	}

	nums1 = []int{2, 3}
	nums2 = []int{0, 5}
	m = findMedianSortedArrays(nums1, nums2)
	if m != 2.5 {
		t.Errorf("Expect 2.5, got=%f", m)
	}

	nums1 = []int{1, 1}
	nums2 = []int{2, 3}
	m = findMedianSortedArrays(nums1, nums2)
	if m != 1.5 {
		t.Errorf("Expect 1.5, got=%f", m)
	}

}
