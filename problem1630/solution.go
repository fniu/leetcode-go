// https://leetcode.com/problems/arithmetic-subarrays/
//

package problem1630

import (
	"sort"
)

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	res := []bool{}
	length := len(l)

Exit:
	for i := 0; i < length; i++ {
		li, ri := l[i], r[i]
		if ri-li <= 1 {
			res = append(res, true)
			continue
		}
		var subnums = make([]int, ri-li+1)
		copy(subnums, nums[li:ri+1])
		sort.Ints(subnums)
		var diff = subnums[1] - subnums[0]
		for si := range subnums {
			if si == 0 {
				continue
			}
			if diff != (subnums[si] - subnums[si-1]) {
				res = append(res, false)
				continue Exit
			}
		}
		res = append(res, true)
	}

	return res

}
