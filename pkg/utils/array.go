package utils

import (
	"golang.org/x/exp/constraints"
)

func Contains[T constraints.Ordered](needle T, haystack []T) bool {
	for _, n := range haystack {
		if needle == n {
			return true
		}
	}
	return false
}
