// https://leetcode.com/problems/subarray-with-elements-greater-than-varying-threshold

package problem2334

func validSubarraySize(nums []int, threshold int) int {

	st := []int{}
	nums = append(nums, 0)
	for i := range nums {
		if nums[i] > threshold {
			return 1
		}
		for len(st) > 0 && nums[i] < nums[st[len(st)-1]] {
			val := nums[st[len(st)-1]]
			st = st[:len(st)-1]
			j := -1
			if len(st) > 0 {
				j = st[len(st)-1]
			}
			if val > threshold/(i-j-1) {
				return i - j - 1
			}
		}
		st = append(st, i)
	}
	return -1

}

func validSubarraySize2(nums []int, threshold int) int {
	var stack []int = []int{}

	for i := 0; i < len(nums); i++ {
		if nums[i] > threshold {
			return 1
		}

		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			var area, width int = 0, 0

			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) > 0 {
				width = i - 1 - stack[len(stack)-1]
				area = nums[idx] * width
			} else {
				width = i
				area = nums[idx] * width
			}
			if area > threshold {
				return width
			}
		}

		stack = append(stack, i)
	}

	for i := 0; i < len(stack); i++ {
		var width int = 0
		if i == 0 {
			width = len(nums)
		} else {
			width = len(nums) - stack[i]
		}
		if nums[stack[i]]*width > threshold {
			return width
		}
	}

	return -1
}

func validSubarraySize3(nums []int, threshold int) int {
	//every element > threshold /k
	stack := make([]int, 0)
	nums = append(nums, 0)
	n := len(nums)
	rs := -1
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[i] <= nums[stack[len(stack)-1]] {
			curMin := nums[stack[len(stack)-1]]
			right := i - 1
			left := -1
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}
			curLen := right - left

			if curMin >= threshold/curLen+1 {
				rs = curLen
			}
		}
		stack = append(stack, i)
	}

	return rs
}
