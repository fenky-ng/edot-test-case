package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(input string) (output string) {
	hash := sha256.New()
	hash.Write([]byte(input))
	output = hex.EncodeToString(hash.Sum(nil))
	return
}
