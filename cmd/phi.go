package main

import (
	"fmt"

	primenumbers "github.com/marksaravi/encryption-fun-go/prime-numbers"
)

func main() {

	const N = 36
	nCoPrimes := make([]int, 0, N)
	// const P = 61
	// const Q = 53

	for n := 1; n <= N; n++ {
		coprimes := primenumbers.FindCoprimes(n)
		nCoPrimes = append(nCoPrimes, len(coprimes))
		fmt.Printf("%3d", n)
	}
	fmt.Println()
	for i := 0; i < N; i++ {
		fmt.Printf("%3d", nCoPrimes[i])
	}
	fmt.Println()
}
