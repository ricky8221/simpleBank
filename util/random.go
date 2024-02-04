package uti

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generator
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generator
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generator
func randomOwner() string {
	return RandomString(6)
}

// RandomMoney generator
func RandomMoney() int64 {
	return RandomInt(0, 10000)
}

// RandomCurrency generator
func RandomCurrency() string {
	currencies := []string{"USD", "RMB", "EUR", "GBP", "JPY", "KRW", "HKD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
