package main

import (
	"fmt"
)

func xorEncryptDecrypt(input string, key byte) string {
	data := []byte(input)
	for i := 0; i < len(data); i++ {
		data[i] ^= key
	}
	return string(data)
}

func main() {
	originalText := "Sonunu düşünen kahraman olamaz."
	key := byte(7)

	encryptedText := xorEncryptDecrypt(originalText, key)
	fmt.Println("Şifrelenmiş metin:", encryptedText)

	decryptedText := xorEncryptDecrypt(encryptedText, key)
	fmt.Println("Çözülmüş metin:", decryptedText)
}
