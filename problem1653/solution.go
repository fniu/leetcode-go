// https://leetcode.com/problems/minimum-deletions-to-make-string-balanced/
// DP
// Time O(n), Space O(n)

package problem1653

func minimumDeletions(s string) int {
	var n, bCount, removeCost int = len(s), 0, 0
	for i := 0; i < n; i++ {
		if s[i] == 'a' {
			removeCost = min(removeCost+1, bCount)
		} else {
			bCount++
		}
	}
	return removeCost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
