package str

import "testing"

func TestContains(t *testing.T) {
	t.Log(Contains("123", []string{"123", "456"}))
	t.Log(Contains("123", []string{"1234", "456"}))
}
