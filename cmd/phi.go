package main

import (
	"fmt"

	"github.com/marksaravi/encryption-fun-go/mathematics"
)

func main() {
	fmt.Println(mathematics.D(17, 780))
	fmt.Println(mathematics.CarmichaelOfPQ(61, 53))
	fmt.Println()
	const N = 36
	nCoPrimes := make([]int, 0, N)
	carmichaels := make([]int, 0, N)
	// const P = 61
	// const Q = 53

	for n := 1; n <= N; n++ {
		coprimes := mathematics.FindCoprimes(n)
		nCoPrimes = append(nCoPrimes, len(coprimes))
		c := mathematics.FindCarmichael(n)
		carmichaels = append(carmichaels, c)
		fmt.Printf("%3d", n)
	}
	fmt.Println()
	for i := 0; i < N; i++ {
		fmt.Printf("%3d", carmichaels[i])
	}
	fmt.Println()
	for i := 0; i < N; i++ {
		fmt.Printf("%3d", nCoPrimes[i])
	}
	fmt.Println()
}
