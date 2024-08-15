package helper

import (
	"crypto/rand"
	"math/big"
)

const shareCodeChars = "ABCDEFGHJKMNPQRSTUVXYZ23456789"
const shareCodeCharsLen = int64(len(shareCodeChars))

func GenerateRandomCode() string {
	length := big.NewInt(shareCodeCharsLen)
	b := make([]byte, 6)

	for i := range b {
		n, err := rand.Int(rand.Reader, length)
		if err != nil {
			b[i] = shareCodeChars[i]
		} else {
			b[i] = shareCodeChars[n.Int64()]
		}
	}
	return string(b)
}
