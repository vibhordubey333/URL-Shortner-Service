package service

import (
	"math/rand"
)

const (
	generatedKeySize = 6
	chars            = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	charsLen         = len(chars)
)

func KeyGenerator() string {
	b := make([]byte, generatedKeySize)
	for i := range b {
		b[i] = chars[rand.Intn(charsLen)]
	}
	return string(b)
}
