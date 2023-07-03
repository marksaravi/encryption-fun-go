package main

import (
	"fmt"

	"github.com/marksaravi/encryption-fun-go/rsa"
)

func main() {
	n, publicKey, PrivateKey := rsa.GenerateKeys()
	fmt.Println(n, publicKey, PrivateKey)
}
