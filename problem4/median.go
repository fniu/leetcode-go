package problem4

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	l1, l2 := len(nums1), len(nums2)

	m := (l1 + l2) / 2
	i := 0
	i1, i2 := 0, 0
	evenTL := (l1+l2)%2 == 0
	prevQ, CurrQ := -1000000, -1000000
	for {
		if (i1 < l1 && i2 < l2 && nums1[i1] <= nums2[i2]) || (i1 < l1 && i2 >= l2) {
			prevQ = CurrQ
			CurrQ = nums1[i1]
			i1++
		} else if (i1 < l1 && i2 < l2 && nums1[i1] > nums2[i2]) || (i1 >= l1 && i2 < l2) {
			prevQ = CurrQ
			CurrQ = nums2[i2]
			i2++
		}
		if i >= m {
			if evenTL {
				return float64(CurrQ+prevQ) / 2
			}
			return float64(CurrQ)
		}
		i++
	}

}
