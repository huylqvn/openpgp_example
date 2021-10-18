package main

import (
	"fmt"

	"github.com/jchavannes/go-pgp/pgp"
)

func main() {
	fmt.Println("Entcrypt test: START")
	pubEntity, _ := pgp.GetEntity([]byte(TestPublicKey), []byte{})
	encrypted, err := pgp.Encrypt(pubEntity, []byte(TestMessage))
	if err != nil {
		return
	}
	fmt.Println("Entcrypt ", string(encrypted))

	privEntity, _ := pgp.GetEntity([]byte(TestPublicKey), []byte(TestPrivateKey))
	decrypted, err := pgp.Decrypt(privEntity, encrypted)
	if err != nil {
		return
	}
	fmt.Println("Decrypted ", string(decrypted))

	decryptedMessage := string(decrypted)
	if decryptedMessage == TestMessage {
		fmt.Println("Decrypted message equals original message.")
	}
}

const TestMessage = "hello world"

const TestPrivateKey = `-----BEGIN PGP PRIVATE KEY BLOCK-----

-----END PGP PRIVATE KEY BLOCK-----`

const TestPublicKey = `-----BEGIN PGP PUBLIC KEY BLOCK-----

-----END PGP PUBLIC KEY BLOCK-----`
