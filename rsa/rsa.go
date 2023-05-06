package rsa

import primenumbers "github.com/marksaravi/encryption-fun-go/prime-numbers"

func generateKeys() int64 {
	P, Q := primenumbers.GetPrimesPQ()
	var N int64 = P * Q
	// n, e := primenumbers.NE(p, q)

	return N
}
