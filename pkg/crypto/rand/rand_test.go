package rand

import "testing"

func TestInt(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(Int(1000))
	}
}

func TestIntString(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(IntString(10))
	}
}

func TestString(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(String(8))
	}
}
