// https://leetcode.com/problems/find-the-duplicate-number/description/

package problem287

// Floyd's Tortoise and Hare (Cycle Detection)
func findDuplicate(nums []int) int {
	n := len(nums)

	if n == 2 {
		return nums[0]
	}
	slow, fast := nums[0], nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return fast
}

// Binary search
func findDuplicate2(nums []int) int {
	low := 1
	high := len(nums) - 1
	cur := 0
	count := 0
	for low <= high {
		cur = (high + low) / 2
		count = 0
		for _, val := range nums {
			if val <= cur {
				count++
			}
		}
		if count > cur {
			high = cur - 1
		} else {
			low = cur + 1
		}
	}
	return high + 1
}
