package helpers

import (
	"crypto/sha1"
	"fmt"
	"math"
)

func GetHash(key string, m int) int64 {
	h := sha1.New()
	h.Write([]byte(key))
	hash := h.Sum(nil)
	hashStr := fmt.Sprintf("%x", hash)

	sum := int64(0)

	for _, char := range hashStr {
		asciiValue := int64(char)
		sum += asciiValue
	}

	return sum%int64(math.Pow(2.0, float64(m)))
}
