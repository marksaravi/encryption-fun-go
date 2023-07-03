package rsa

import (
	"fmt"

	"github.com/marksaravi/encryption-fun-go/mathematics"
)

func GenerateKeys() (n, publicKey, PrivateKey int) {
	// P, Q := primenumbers.GetPrimesPQ()
	P := 61
	Q := 53
	var N int = P * Q
	carmichael := mathematics.CarmichaelOfPQ(P, Q)

	fmt.Println(carmichael)
	e := 17
	modularMultiplicativeInverse, _, _, _ := mathematics.ModularMultiplicativeInverseGCD1(17, carmichael)

	return N, e, modularMultiplicativeInverse
}
