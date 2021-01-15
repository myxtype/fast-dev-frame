package rand

import (
	"crypto/rand"
	"math/big"
)

var (
	strings = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	numbers = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
)

func Int(n int64) int64 {
	r, _ := rand.Int(rand.Reader, big.NewInt(n))
	return r.Int64()
}

func IntString(n int) string {
	var container string
	length := big.NewInt(10)
	for i := 0; i < n; i++ {
		random, _ := rand.Int(rand.Reader, length)
		container = container + numbers[random.Int64()]
	}
	return container
}

func String(n int) string {
	var container string
	length := big.NewInt(16)
	for i := 0; i < n; i++ {
		random, _ := rand.Int(rand.Reader, length)
		container = container + strings[random.Int64()]
	}
	return container
}
