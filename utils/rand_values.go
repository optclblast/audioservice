package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func randomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomLogin() string {
	return randomString(15)
}

func RandomPassword() string {
	return randomString(25)
}

func RandomID() int64 {
	return RandomInt(1, 10000)
}

func RandomNumber(from int64, to int64) int64 {
	return RandomInt(from, to)
}
