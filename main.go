package main

import "fmt"

func Encrypt(key int, plainText string) (result string) {

}

func Decrypt(key int, CipherText string) (result string) {

}

func main() {
	plainText := "ANJAS"
	fmt.Println("Plain Text: ", plainText)

	encrypted := Encrypt(5, plainText)
	fmt.Println("Encrypted text: ", encrypted)

	decrypted := Decrypt(5, encrypted)
	fmt.Println("Decrypted text: ", decrypted)
}