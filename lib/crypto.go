package lib

import (
    "fmt"
    "io"

    "crypto/aes"
    "crypto/cipher"
	"crypto/sha256"
	"crypto/rand"

    "encoding/hex"
)

func Encrypt(text string, password string) string {

    plain := []byte(text)
    data  := []byte(password)
    key   := sha256.Sum256(data)

    // Create a new cipher block from the key
    block, err := aes.NewCipher(key[:])
    if err != nil {
        panic(err.Error())
    }

    // Create a new GCM
    // https://en.wikipedia.org/wiki/Galois/Counter_Mode
    // https://golang.org/pkg/crypto/cipher/#NewGCM
    aesGCM, err := cipher.NewGCM(block)
    if err != nil {
        panic(err.Error())
    }

    // Create a nonce. Nonce should be from GCM
    nonce := make([]byte, aesGCM.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        panic(err.Error())
    }

    // Encrypt the data using aesGCM.Seal. Since we don't want to save the nonce
    // somewhere else in this case, we add it as a prefix to the encrypted data.
    // The first nonce argument in Seal is the prefix.
    ciphered := aesGCM.Seal(nonce, nonce, plain, nil)
    return fmt.Sprintf("%x", ciphered)
}

func Decrypt(text string, password string) string {
    data   := []byte(password)
    key    := sha256.Sum256(data)
    enc, _ := hex.DecodeString(text)

    // Create a new cipher block from the key
    block, err := aes.NewCipher(key[:])
    if err != nil {
        panic(err.Error())
    }

    // Create a new GCM
    aesGCM, err := cipher.NewGCM(block)
    if err != nil {
        panic(err.Error())
    }

    // Get the nonce size
    nonceSize := aesGCM.NonceSize()

    // Extract the nonce from the encrypted data
    nonce, ciphered := enc[:nonceSize], enc[nonceSize:]

    // Decrypt the data
    plain, err := aesGCM.Open(nil, nonce, ciphered, nil)
    if err != nil {
        panic(err.Error())
    }

    return fmt.Sprintf("%s", plain)
}
