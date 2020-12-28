package problem1249

import "testing"

func TestSolution(t *testing.T) {

	expectedValidStr(t, "(abc)", minRemoveToMakeValid("(abc)"))
	expectedValidStr(t, "", minRemoveToMakeValid("))(("))
	expectedValidStr(t, "ab(c)", minRemoveToMakeValid("(ab(c)"))
	expectedValidStr(t, "a(b(c)d)", minRemoveToMakeValid("(a(b(c)d)"))
	expectedValidStr(t, "ab(c)d", minRemoveToMakeValid("a)b(c)d"))

}

func expectedValidStr(t *testing.T, expected, got string) {
	t.Helper()
	if expected != got {
		t.Errorf("Expected %q, got=%q", expected, got)
	}
}
