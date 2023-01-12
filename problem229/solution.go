// https://leetcode.com/problems/majority-element-ii/
// Time O(n)
// Space O(1)

package problem229

func majorityElement(nums []int) []int {
	ln := len(nums)

	if ln < 2 {
		return nums
	}

	rest := make([]int, 0, 2)
	n := ln/3 + 1
	counts := map[int]int{}
	for _, num := range nums {
		if counts[num] < 0 {
			continue
		}
		counts[num]++
		if counts[num] >= n {
			rest = append(rest, num)
			if len(rest) >= 2 {
				return rest
			}
			counts[num] = -1
		}
	}

	return rest

}
