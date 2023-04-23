package main

import (
	"fmt"

	primenumbers "github.com/marksaravi/encryption-fun-go/prime-numbers"
)

func main() {
	for i := 1; i < 2; i++ {
		fmt.Println(primenumbers.Phi(i))
	}
}
