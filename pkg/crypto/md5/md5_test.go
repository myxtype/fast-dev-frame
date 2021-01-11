package md5

import "testing"

func TestEncryptString(t *testing.T) {
	h, err := EncryptString("")
	if err != nil {
		t.Error(err)
	}
	t.Log(h)
}
