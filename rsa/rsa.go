package rsa

import (
	"time"

	"github.com/marksaravi/encryption-fun-go/mathematics"
)

func GenerateKeys() (n, publicKey, PrivateKey int64) {
	// P, Q := primenumbers.GetPrimesPQ()
	P := int64(61)
	Q := int64(53)
	var N int64 = P * Q
	carmichael := mathematics.CarmichaelOfPQ(P, Q)

	coPrimes := mathematics.FindCoprimes(carmichael)
	i := time.Now().Nanosecond() % len(coPrimes)
	e := coPrimes[i]
	// e := int64(17)
	modularMultiplicativeInverse, _, _, _ := mathematics.ModularMultiplicativeInverseGCD1(17, carmichael)

	return N, e, modularMultiplicativeInverse
}

func Encrypt(a, PublicKey, N int64) int64 {
	return mathematics.ReminderOfPower(a, PublicKey, N)
}

func Decrypt(a, PrivateKey, N int64) int64 {
	return mathematics.ReminderOfPower(a, PrivateKey, N)
}
