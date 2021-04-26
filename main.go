package main

import (
	"crypto/sha1"
	"fmt"
	"math"
)

func getHash(key string, m int) int {
	h := sha1.New()
	h.Write([]byte(key))
	hash := h.Sum(nil)
	hashStr := fmt.Sprintf("%x", hash)

	sum := 0

	for _, char := range hashStr {
		asciiValue := int(char)
		sum += asciiValue
	}

	return sum%int(math.Pow(2.0, float64(m)))
}

func main() {
	testingKey := "rao"
	h := sha1.New()
	h.Write([]byte(testingKey))
	hash := h.Sum(nil)
	hashStr := fmt.Sprintf("%x", hash)

	sum := 0

	for _, char := range hashStr {
		asciiValue := int(char)
		sum += asciiValue
	}

	m := 10.0

	fmt.Println(sum%int((math.Pow(2.0, m))))
}