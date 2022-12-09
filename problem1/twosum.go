package problem1

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)
	for i, a := range nums {
		if val, ok := m[a]; ok {
			return []int{i, val}
		}
		m[target-a] = i
	}
	panic("No answer can be found.")
}
