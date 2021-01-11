package aes

import "testing"

func TestEncryptToHex(t *testing.T) {
	key := []byte("2qagBhx4gh6zF0sK")
	data := []byte("123456789abcdef")

	d, err := EncryptToHex(data, key)
	if err != nil {
		t.Error(err)
	}

	t.Log(d)

	e, err := DecryptFromHex(d, key)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(e))
}
