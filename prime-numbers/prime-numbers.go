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
	i1 := randomGenerator.Intn(len(PrimeNumbers) / 2)
	i2 := randomGenerator.Intn(len(PrimeNumbers) / 2)
	for i1 == i2 {
		i2 = randomGenerator.Intn(len(PrimeNumbers))
	}
	P = int64(PrimeNumbers[i1])
	Q = int64(PrimeNumbers[i2])
	return
}

func FindCoprimes(a int) []int {
	coprimes := make([]int, 0, 1000)
	coprimes = append(coprimes, 1)
	for n := 2; n < a; n++ {
		iscoprime := true
		for p := 0; PrimeNumbers[p] <= n; p++ {
			if a%PrimeNumbers[p] == 0 && n%PrimeNumbers[p] == 0 {
				iscoprime = false
				break
			}
		}
		if iscoprime {
			coprimes = append(coprimes, n)
		}
	}
	return coprimes
}
