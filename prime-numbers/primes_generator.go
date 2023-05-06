package primenumbers

import (
	"fmt"
	"os"
)

func GeneratePrimeNumbers(N int) {
	// const N = 1000000
	pn := make([]int, N)
	pn[0] = 2
	pn[1] = 3
	const NPL = 20
	f, _ := os.Create("./primes.txt")
	numberPerLine := 0
	primeCounter := 2
	lastTest := pn[1]
	for primeCounter < N {
		isprime := true
		lastTest += 2
		for j := 0; j < primeCounter; j++ {
			if lastTest/2 <= pn[j] {
				break
			}
			if lastTest%pn[j] == 0 {
				isprime = false
				break
			}
		}
		if isprime {
			pn[primeCounter] = lastTest
			primeCounter++
			f.WriteString(fmt.Sprintf("%d,", lastTest))
			numberPerLine++
			if numberPerLine%NPL == 0 {
				f.WriteString("\n")
				numberPerLine = 0
			}
		}

	}
	f.Close()
}
