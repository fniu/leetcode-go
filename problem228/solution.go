// https://leetcode.com/problems/summary-ranges/
// time: O(n)
// space: O(n)

package problem228

import "fmt"

func summaryRanges(nums []int) []string {

	result := []string{}
	if len(nums) == 0 {
		return result
	}
	result = append(result, fmt.Sprintf("%d", nums[0]))
	if len(nums) == 1 {
		return result
	}
	val, c := 0, 0
	for i := 1; i < len(nums); i++ {
		val = nums[i]
		if val-nums[i-1] == 1 {
			continue
		}
		if i-c == 1 {
			result = append(result, fmt.Sprintf("%d", val))
			c = i
		} else {
			result[len(result)-1] = result[len(result)-1] + "->" + fmt.Sprintf("%d", nums[i-1])
			result = append(result, fmt.Sprintf("%d", val))
			c = i
		}
	}
	if c != len(nums)-1 {
		result[len(result)-1] = result[len(result)-1] + "->" + fmt.Sprintf("%d", nums[len(nums)-1])
	}
	return result
}
