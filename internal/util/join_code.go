package util

import (
	"crypto/rand"
	"math/big"
)

const (
	letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits  = "0123456789"
)

func GenerateJoinCode() string {
	b := make([]byte, 6)

	for i := 0; i < 3; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		b[i] = letters[num.Int64()]
	}

	for i := 3; i < 6; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		b[i] = digits[num.Int64()]
	}

	return string(b)
}


