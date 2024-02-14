package util

import (
	"fmt"
	"github.com/go-playground/locales/currency"
	"math/rand"
	"strings"
	"time"
	"unicode"
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
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		if i == 0 {
			c = byte(unicode.ToUpper(rune(c)))
		}
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generator
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generator
func RandomMoney() int64 {
	return RandomInt(0, 10000)
}

// RandomCurrency generator
func RandomCurrency() string {
	currencies := []string{string(currency.USD), EUR, CAD}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
