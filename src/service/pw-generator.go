package service

import (
	"log"
	"math/rand"
	"time"
)

const (
	// LowerLetters is the list of lowercase letters.
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"

	// UpperLetters is the list of uppercase letters.
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Digits is the list of permitted digits.
	Digits = "0123456789"

	// Symbols is the list of symbols.
	Symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

func GeneratePassword(numberOfCharacters int, useDigits bool, useSymbols bool) (string, error) {
	characterPool := LowerLetters + UpperLetters
	if useDigits {
		characterPool += Digits
	}
	if useSymbols {
		characterPool += Symbols
	}

	resultString := ""
	for i := 0; i < numberOfCharacters; i++ {
		char, err := getRandomElement(characterPool)
		if err != nil {
			log.Print("Error retrieving char from string")
		}
		resultString += char
	}
	return resultString, nil
}

func getRandomElement(pool string) (string, error) {
	inRune := []rune(pool)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s) // initialize local pseudorandom generator
	randomIndex := r.Intn(len(pool))
	pick := inRune[randomIndex]
	return string(pick), nil
}
