package main

import (
    "fmt"
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "encoding/hex"
)

func main() {
    password := "123456789-123456789-123456789-12"
    iv := "123456789-123456"
    text := "abcdefghijklmnopqrstuvwxyzABCDEF"
    res := Ase256(text, password, iv, aes.BlockSize)
    fmt.Printf("Result: %v\n", res)
}

func Ase256(text string, password string, iv string, blockSize int) string {
    bKey := []byte(password)
    bIV := []byte(iv)
    bPlaintext := PKCS5Padding([]byte(text), blockSize, len(text))
    block, _ := aes.NewCipher(bKey)
    ciphertext := make([]byte, len(bPlaintext))
    mode := cipher.NewCBCEncrypter(block, bIV)
    mode.CryptBlocks(ciphertext, bPlaintext)
    return hex.EncodeToString(ciphertext)
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
    padding := (blockSize - len(ciphertext)%blockSize)
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}
