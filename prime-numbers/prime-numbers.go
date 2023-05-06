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

func FindCoprimes(a int) []int {
	coprimes := make([]int, 0, 1000)
	for n := 2; n < a; n++ {

	}
	return coprimes
}
