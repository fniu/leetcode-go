// https://leetcode.com/problems/number-of-ways-where-square-of-number-is-equal-to-product-of-two-numbers/

package problem1577

func numTriplets(nums1 []int, nums2 []int) int {
	if len(nums1) == 1 && len(nums2) == 1 {
		return 0
	}

	result := 0
	squarei := 0
	for i := range nums1 {
		squarei = nums1[i] * nums1[i]
		result += walkNums(squarei, &nums2)
	}

	for i := range nums2 {
		squarei = nums2[i] * nums2[i]
		result += walkNums(squarei, &nums1)
	}
	return result
}

func walkNums(squarei int, nums *[]int) (result int) {
	j, k := 0, 1
	result = 0
	lenNums := len(*nums)
	for k < lenNums {
		if (*nums)[j]*(*nums)[k] == squarei {
			result += 1
		}
		k += 1
		if k == lenNums {
			j += 1
			k = j + 1
		}
		if j == lenNums-1 {
			break
		}
	}
	return
}
