package primenumbers

import (
	"math/rand"
	"time"
)

var randomSource rand.Source
var randomGenerator *rand.Rand

func init() {
	randomSource := rand.NewSource(time.Now().UnixNano())
	randomGenerator = rand.New(randomSource)
}

func GetPrimesPQ() (P int64, Q int64) {
	i1 := randomGenerator.Intn(len(PrimeNumbers4Digits))
	i2 := randomGenerator.Intn(len(PrimeNumbers4Digits))
	for i1 == i2 {
		i2 = randomGenerator.Intn(len(PrimeNumbers4Digits))
	}
	P = int64(PrimeNumbers4Digits[i1])
	Q = int64(PrimeNumbers4Digits[i2])
	return
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

// A simple method to evaluate Euler Totient Function
func Phi(n int) int {
	var result int = 1
	for i := 2; i < n; i++ {
		if gcd(i, n) == 1 {
			result++
		}
	}
	return result
}

func NE(p, q int) (int, int) {
	n := p * q
	phi := Phi(n)
	e := randomGenerator.Intn(phi)
	for e <= 1 {
		e = randomGenerator.Intn(phi)
	}
	return phi, e
}
