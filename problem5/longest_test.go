package problem5

import (
	"testing"
)

func TestSolution(t *testing.T) {
	s := longestPalindrome("abc")
	expectedString(t, "a", s)

	s = longestPalindrome("bb")
	expectedString(t, "bb", s)

	s = longestPalindrome("b")
	expectedString(t, "b", s)

	s = longestPalindrome("12")
	expectedString(t, "1", s)

	s = longestPalindrome("abcb")
	expectedString(t, "bcb", s)

	s = longestPalindrome("bcb1")
	expectedString(t, "bcb", s)

	s = longestPalindrome("abcbe")
	expectedString(t, "bcb", s)

	s = longestPalindrome("abcb2")
	expectedString(t, "bcb", s)

	s = longestPalindrome("b1234321b")
	expectedString(t, "b1234321b", s)

	s = longestPalindrome("ccb1234321b")
	expectedString(t, "b1234321b", s)

	s = longestPalindrome("cccabcdcba")
	expectedString(t, "abcdcba", s)

	s = longestPalindrome("abcd")
	expectedString(t, "a", s)

	s = longestPalindrome("aaaa")
	expectedString(t, "aaaa", s)

	s = longestPalindrome("1111")
	expectedString(t, "1111", s)

	s = longestPalindrome("1221")
	expectedString(t, "1221", s)

	s = longestPalindrome("1221a1a")
	expectedString(t, "1221", s)

	s = longestPalindrome("12212")
	expectedString(t, "1221", s)
	s = longestPalindrome("21221")
	expectedString(t, "1221", s)

}

func expectedString(t *testing.T, expected, got string) {
	t.Helper()
	if expected != got {
		t.Errorf("Expect '%q', got=%q", expected, got)
	}
}
