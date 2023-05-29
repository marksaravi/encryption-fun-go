package primenumbers

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/marksaravi/encryption-fun-go/mathematics"
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
		if mathematics.GreatestCommonDivisor(a, n) == 1 {
			coprimes = append(coprimes, n)
		}
		// iscoprime := true
		// for p := 0; PrimeNumbers[p] <= n; p++ {
		// 	if a%PrimeNumbers[p] == 0 && n%PrimeNumbers[p] == 0 {
		// 		iscoprime = false
		// 		break
		// 	}
		// }
		// if iscoprime {
		// 	coprimes = append(coprimes, n)
		// }
	}
	return coprimes
}

func FindLCM(a, b int) int {
	k := 1
	fmt.Println(a, b)
	for i := 2; i < a && i < b; i++ {
		if a%i == 0 && b%i == 0 {
			k = i
		}
	}

	return a * b / k
}
