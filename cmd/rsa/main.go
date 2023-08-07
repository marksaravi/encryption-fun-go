package main

import (
	"fmt"

	"github.com/marksaravi/encryption-fun-go/rsa"
)

func main() {
	n, publicKey, privateKey := rsa.GenerateKeys()
	fmt.Printf("N: %d\nPublic Key: %d\nPrivate Key: %d\n", n, publicKey, privateKey)
	a := int64(765)
	fmt.Printf("Number to Encrypt: %d\n", a)
	encrypted := rsa.Encrypt(a, publicKey, n)
	fmt.Printf("Encrypted: %d\n", encrypted)
	decrypted := rsa.Encrypt(encrypted, privateKey, n)
	fmt.Printf("Decrypt: %d\n", decrypted)
}
