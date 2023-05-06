package rsa

import primenumbers "github.com/marksaravi/encryption-fun-go/prime-numbers"

func generateKeys() {
	// p, q := primenumbers.GetPrimes()
	// n, e := primenumbers.NE(p, q)

	// return
}

func genPQ() (int, int) {
	p, q := primenumbers.GetPrimes()

	return p, q
}

func genPublicKey(p, q int) (int, int) {
	n, e := primenumbers.NE(p, q)

	return n, e
}

func genPrivateKey(p, q, e int) (int, int) {
	// phi := (p - 1) * (q - 1)
	// x := float64(phi+1) / float64(e)
	// // if x<1
	// return n, e
	return 0, 0
}
