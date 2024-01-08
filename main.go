package main

import (
	"fmt"
	"strings"
)

const originalLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func hashLetter(key int, letter string) (result string) {
	runes := []rune(letter)

	lastLetters := string(runes[len(letter)-key : len(letter)])

	restLetters := string(runes[0 : len(letter)-key])

	return fmt.Sprintf(`%s%s`, lastLetters, restLetters)
}

func Encrypt(key int, plainText string) (result string) {
	hashLetter := hashLetter(key, originalLetter)

	var hashedString = ""

	findOne := func(r rune) rune {
		position := strings.Index(originalLetter, string([]rune{r}))

		if position != -1 {
			letterPosition := (position + len(originalLetter)) % len(originalLetter)

			hashedString = hashedString + string(hashLetter[letterPosition])

			return r
		}

		return r
	}

	strings.Map(findOne, plainText)

	return hashedString
}

func Decrypt(key int, CipherText string) (result string) {
	hashLetter := hashLetter(key, originalLetter)

	var hashedString = ""

	findOne := func(r rune) rune {
		position := strings.Index(hashLetter, string([]rune{r}))

		if position != -1 {
			letterPosition := (position + len(originalLetter)) % len(originalLetter)

			hashedString = hashedString + string(originalLetter[letterPosition])

			return r
		}

		return r
	}

	strings.Map(findOne, CipherText)

	return hashedString
}

func main() {
	plainText := "ANJASFEDO"
	fmt.Println("Plain Text: ", plainText)

	encrypted := Encrypt(5, plainText)
	fmt.Println("Encrypted text: ", encrypted)

	decrypted := Decrypt(5, encrypted)
	fmt.Println("Decrypted text: ", decrypted)
}
