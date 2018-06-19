package prob

// Given a string, find the longest palindromic contiguous substring. If there
// are more than one with the maximum length, return any one.
//
// For example, the longest palindromic substring of "aabcdcb" is "bcdcb". The
// longest palindromic substring of "bananas" is "anana".

func LongestPalindrome(s string) string {
	result := ""
	for i := range s {
		for j := i + 1; j <= len(s); j++ {
			candidate := s[i:j]
			if IsPalindrome(candidate) && len(candidate) > len(result) {
				result = candidate
			}
		}
	}
	return result
}

func IsPalindrome(s string) bool {
	// TODO: This should really compare runes.
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
