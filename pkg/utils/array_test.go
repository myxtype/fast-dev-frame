package utils

import "testing"

func TestContains(t *testing.T) {
	t.Log(Contains("123", []string{"213", "123", "456"}))
	t.Log(Contains("123", []string{"1234", "456"}))
	t.Log(Contains("123", []string{"1234", "456"}))
	t.Log(Contains("v", []string{"1", "3", "-1", "23", "32", "f", "v"}))
}
