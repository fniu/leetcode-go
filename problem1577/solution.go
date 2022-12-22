// https://leetcode.com/problems/number-of-ways-where-square-of-number-is-equal-to-product-of-two-numbers/

package problem1577

func numTriplets(nums1 []int, nums2 []int) int {
	if len(nums1) == 1 && len(nums2) == 1 {
		return 0
	}
	cnt1, cnt2 := freqMap(nums1), freqMap(nums2)

	return calculate(nums1, nums2, cnt1, cnt2) + calculate(nums2, nums1, cnt2, cnt1)
}

func freqMap(nums []int) *map[int]int {
	counts := map[int]int{}
	for i := range nums {
		counts[nums[i]] += 1
	}
	return &counts
}

func calculate(nums1, nums2 []int, cnt1, cnt2 *map[int]int) int {
	var result int
	var xSqr, acc int
	for x1, c := range *cnt1 {
		acc = 0
		xSqr = x1 * x1
		for _, x2 := range nums2 {
			y := xSqr / x2
			if xSqr%x2 != 0 {
				continue
			}
			if val, ok := (*cnt2)[y]; ok {
				acc += val
				if y == x2 {
					acc -= 1
				}
			}
		}
		result += c * acc / 2
	}
	return result
}

func numTriplets2(nums1 []int, nums2 []int) int {
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
