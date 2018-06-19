package prob

import "strconv"

// This problem was asked by Facebook.
//
// Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the
// number of ways it can be decoded.
//
// For example, the message '111' would give 3, since it could be decoded as
// 'aaa, 'ka', and 'ak'.

func NumDecodes(s string) int {
	result := 0
	if isValid(s) {
		result += 1
	}
	if len(s) > 1 && isValid(s[:1]) {
		result += NumDecodes(s[1:])
	}
	if len(s) > 2 && isValid(s[:2]) {
		result += NumDecodes(s[2:])
	}
	return result
}

func NumDecodes2(s string) int {
	var rec func(string, string) int
	rec = func(head, tail string) (result int) {
		// Base case: return 1 if the head is in (1,26), or zero otherwise.
		if len(tail) == 0 {
			if isValid(head) {
				result += 1
			}
			return
		}

		// Recursive case: try pulling 1 or 2 characters off the front of the
		// string. If that forms a valid number in (1,26), then recurse and
		// repeat on the tail.
		h, t := tail[:1], tail[1:]
		if isValid(h) {
			result += rec(h, t)
		}
		if len(tail) >= 2 {
			h, t = tail[:2], tail[2:]
			if isValid(h) {
				result += rec(h, t)
			}
		}
		return
	}
	return rec("", s)
}

func isValid(s string) bool {
	n, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return n >= 1 && n <= 26
}
