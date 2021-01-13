package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main() {
	text := os.Args[1]
	pass := os.Args[2]
	enc2 := os.Args[3]

	enc1 := encrypt(text, pass)
	fmt.Printf("enc1 : %s\n", enc1)

	dec1 := decrypt(enc1, pass)
	fmt.Printf("dec1 : %s\n", dec1)

	dec2 := decrypt(enc2, pass)
	fmt.Printf("dec2 : %s\n", dec2)
}

func encrypt(text string, password string) string {

	data := []byte(password)
	key  := sha256.Sum256(data)
	plaintext := []byte(text)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key[:])
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

func decrypt(text string, password string) string {

	data := []byte(password)
	key  := sha256.Sum256(data)

	// key, _ := hex.DecodeString(password)
	enc, _ := hex.DecodeString(text)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key[:])
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}
