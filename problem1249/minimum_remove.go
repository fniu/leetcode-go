package problem1249

import (
	"bytes"
	"strings"
)

func minRemoveToMakeValid(s string) string {
	stack := []int{}
	newS := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == 40 {
			stack = append(stack, i)
			newS = append(newS, 32)
		} else if s[i] == 41 {
			if len(stack) > 0 && stack[len(stack)-1] < len(s) {
				newS = append(newS, 41)
				newS[stack[len(stack)-1]] = 40
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, len(s))
				newS = append(newS, 32)
			}
		} else {
			newS = append(newS, byte(s[i]))
		}
	}
	var vs bytes.Buffer
	for _, val := range newS {
		if val > 39 {
			vs.WriteByte(val)
		}
	}
	return vs.String()
}

func minRemoveToMakeValid2(s string) string {
	stack := []int{}
	indexMap := map[int]bool{}
	for i := 0; i < len(s); i++ {
		if s[i] == 40 {
			stack = append(stack, i+1)
			indexMap[i] = true
		}
		if s[i] == 41 {
			if len(stack) > 0 {
				if stack[len(stack)-1] > 0 {
					indexMap[stack[len(stack)-1]-1] = false
					stack = stack[:len(stack)-1]
					continue
				}
			}
			stack = append(stack, -(i + 1))
			indexMap[i] = true
		}
	}
	var vs bytes.Buffer
	for i, v := range s {
		if val, ok := indexMap[i]; ok && val {
			continue
		}
		vs.WriteString(string(v))
	}
	return vs.String()
}

func minRemoveToMakeValidEasy(s string) string {

	stack := []int{}
	indexMap := map[int]bool{}

	for i := 0; i < len(s); i++ {
		if s[i] == 40 {
			stack = append(stack, i+1)
			indexMap[i] = true
		}
		if s[i] == 41 {
			if len(stack) > 0 {
				if stack[len(stack)-1] > 0 {
					indexMap[stack[len(stack)-1]-1] = false
					stack = stack[:len(stack)-1]
					continue
				}
			}
			stack = append(stack, -(i + 1))
			indexMap[i] = true
		}
	}

	vs := []string{}
	for i, v := range s {
		if val, ok := indexMap[i]; ok && val {
			continue
		}
		vs = append(vs, string(v))
	}

	return strings.Join(vs, "")
}
