package rsa

import (
	"time"

	"github.com/marksaravi/encryption-fun-go/mathematics"
	primenumbers "github.com/marksaravi/encryption-fun-go/prime-numbers"
)

func GenerateKeys() (n, publicKey, PrivateKey int64) {
	P, Q := primenumbers.GetPrimesPQ()
	// P := 61
	// Q := 53
	var N int64 = P * Q
	carmichael := mathematics.CarmichaelOfPQ(P, Q)

	coPrimes := mathematics.FindCoprimes(carmichael)
	i := time.Now().Nanosecond() % len(coPrimes)
	e := coPrimes[i]
	modularMultiplicativeInverse, _, _, _ := mathematics.ModularMultiplicativeInverseGCD1(17, carmichael)

	return N, e, modularMultiplicativeInverse
}
