package problem5

import (
	"math"
)

func longestPalindrome(s string) string {
	var l, la string
	l = string(s[0])
	lens := len(s)
	if lens == 1 {
		return s
	}
	if lens == 2 {
		if s[0] == s[1] {
			return s
		}
		return l
	}
	var i2, end, a, b int
	for i := 0.5; int(math.Round(i)) < lens; i += 0.5 {
		la = ""
		i2 = int(math.Round(i))
		for j := i2 - 1; j >= 0; j-- {
			if i2 > int(i) {
				end = i2 + (i2 - j) - 1
			} else {
				end = i2 + (i2 - j)
			}
			if end >= lens {
				break
			}
			if s[j] != s[end] {
				break
			}
			a, b = j, end
		}
		la = s[a : b+1]
		if len(la) == 0 && s[i2-1] == s[i2] {
			la = s[i2-1 : i2+1]
		}
		la = string(la)
		if len(la) > len(l) {
			l = la
		}
	}
	return l
}

func longestPalindrome2(s string) string {
	var l, la string
	l = string(s[0])
	lens := len(s)
	if lens == 1 {
		return s
	}
	if lens == 2 {
		if s[0] == s[1] {
			return s
		}
		return l
	}
	found := false
	var start, i2, end int

	for i := 0.5; int(math.Round(i)) < lens; i += 0.5 {
		found = false
		la = ""
		i2 = int(math.Round(i))
		for j := 0; j < i2; j++ {

			if !found {
				start = j
			}
			if i2 > int(i) {
				end = i2 - j + i2 - 1
			} else {
				end = i2 - j + i2
			}

			if end >= lens {
				continue
			}

			if s[j] != s[end] {
				found = false
				continue
			}
			found = true
			if j == i2-1 {
				if i2 > int(i) {
					la = s[start : i2-start+i2]
				} else {
					la = s[start : i2-start+i2+1]
				}
			}

		}
		if !found && s[i2-1] == s[i2] {
			la = (s[i2-1 : i2+1])
		}
		la = string(la)
		if len(la) > len(l) {
			l = la
		}
	}

	return l
}
