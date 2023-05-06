package main

import (
	"fmt"

	primenumbers "github.com/marksaravi/encryption-fun-go/prime-numbers"
)

func main() {
	fmt.Printf("4   digits: %d\n", len(primenumbers.PrimeNumbers4Digits))
	// fmt.Printf("5   digits: %d\n", len(primenumbers.PrimeNumbers5Digits))
	// fmt.Printf("6   digits: %d\n", len(primenumbers.PrimeNumbers6Digits))
	// fmt.Printf("7   digits: %d\n", len(primenumbers.PrimeNumbers7Digits))
	// fmt.Printf("8   digits: %d\n", len(primenumbers.PrimeNumbers8Digits))
	// fmt.Printf("all digit: %d\n", len(primenumbers.PrimeNumbersMillion))
}
