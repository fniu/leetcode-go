// https://leetcode.com/problems/minimum-elements-to-add-to-form-a-given-sum/

package problem1785

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	diff := goal - sum
	// set diff non-negative, for easier code
	if diff < 0 {
		diff = -diff
	}
	count := diff / limit
	if diff%limit != 0 {
		count++
	}
	return count
}
