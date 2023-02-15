package util

import (
	cryptoRand "crypto/rand"
	"math/bits"
	"math/rand"
	"time"
	"unsafe"
)

// Previous defined bytes, do not change this.
var (
	DefaultAlphabet   = []byte("QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz")
	DefaultDigit      = []byte("0123456789")
	DefaultAlphaDigit = []byte("QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz0123456789")
	DefaultSymbol     = []byte("QWERTYUIOPLKJHGFDSAZXCVBNMabcdefghijklmnopqrstuvwxyz0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_{|}~`") // nolint: lll
)

func RandNumber(n int) string {
	return randString(n, DefaultDigit)
}

func RandString(n int) string {
	return randString(n, DefaultAlphaDigit)
}

func randString(length int, alphabets []byte) string {
	b := randBytes(length, alphabets)
	return *(*string)(unsafe.Pointer(&b))
}

func randBytes(length int, alphabets []byte) []byte {
	b := make([]byte, length)
	if _, err := cryptoRand.Read(b); err == nil {
		for i, v := range b {
			b[i] = alphabets[v%byte(len(alphabets))]
		}
		return b
	}

	bn := bits.Len(uint(len(alphabets)))
	mask := int64(1)<<bn - 1
	max := 63 / bn
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63() + rand.Int63() + rand.Int63()))

	// A rand.Int63() generates 63 random bits, enough for alphabets letters!
	for i, cache, remain := 0, r.Int63(), max; i < length; {
		if remain == 0 {
			cache, remain = r.Int63(), max
		}
		if idx := int(cache & mask); idx < len(alphabets) {
			b[i] = alphabets[idx]
			i++
		}
		cache >>= bn
		remain--
	}
	return b
}
