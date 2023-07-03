package main

import (
	"fmt"

	"github.com/marksaravi/encryption-fun-go/rsa"
)

func main() {
	n, publicKey, privateKey := rsa.GenerateKeys()
	fmt.Println(n, publicKey, privateKey)
	a := int64(765)
	fmt.Println(a)
	encrypted := rsa.Encrypt(a, publicKey, n)
	fmt.Println(encrypted)
	decrypted := rsa.Encrypt(encrypted, privateKey, n)
	fmt.Println(decrypted)
}
