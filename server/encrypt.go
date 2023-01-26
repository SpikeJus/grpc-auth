package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"log"
)

var keys rsa.PrivateKey

func Init() {
	reader := rand.Reader
	bitSize := 2048
	prKey, _ := rsa.GenerateKey(reader, bitSize)
	keys = *prKey
}

func GetPublicKey() rsa.PublicKey {
	return keys.PublicKey
}

func Decrypt(encrypted []byte) string {
	// decrypted, err := keys.Decrypt(nil, []byte(encrypted), &rsa.OAEPOptions{Hash: crypto.SHA256})
	decrypted, err := rsa.DecryptOAEP(sha256.New(), nil, &keys, []byte(encrypted), nil)
	errDecryption(err)
	return string(decrypted)
}

func errDecryption(err error) {
	if err != nil {
		log.Println("decryption error: ", err)
	}
}
