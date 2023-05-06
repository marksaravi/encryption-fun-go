package main

import (
	"fmt"

	primenumbers "github.com/marksaravi/encryption-fun-go/prime-numbers"
)

func main() {
	for i := 1; i <= 36; i++ {
		coprimes := primenumbers.FindCoprimes(i)
		// fmt.Printf("%d: %d (%v)\n", i, len(coprimes), coprimes)
		fmt.Printf("%2d: %2d\n", i, len(coprimes))
	}
}
