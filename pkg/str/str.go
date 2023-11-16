package str

import "strings"

func Contains(needle string, haystack []string) bool {
	left, right := 0, len(haystack)-1
	for left <= right {
		mid := left + (right-left)/2
		if strings.Compare(haystack[mid], needle) == 0 {
			return true
		} else if strings.Compare(haystack[mid], needle) < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
