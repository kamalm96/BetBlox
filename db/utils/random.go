package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomString(n int) string {

	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomDate(n time.Time) time.Time {
	return time.Date(RandomInt(1950, 2004), time.Month(RandomInt(1, 12)), RandomInt(1, 30), 0, 0, 0, 0, time.UTC)
}

func RandomUsername() string {
	return RandomString(7)
}

func RandomEmail() string {
	return RandomString(6) + "@email.com"
}

func RandomLast4SSN() int {
	return RandomInt(1000, 9999)
}
