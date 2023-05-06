package main

import (
	"fmt"

	primenumbers "github.com/marksaravi/encryption-fun-go/prime-numbers"
)

func main() {

	const P = 61
	const Q = 53

	LCM := primenumbers.FindLCM(P-1, Q-1)
	fmt.Println(LCM)
	coprimes := primenumbers.FindCoprimes(LCM)
	fmt.Printf("%2d: %2d\n", LCM, coprimes)
}
