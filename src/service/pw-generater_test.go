package service

import (
	"fmt"
	"testing"
)

func TestOnlyCharacters(t *testing.T) {
	result, err := GeneratePassword(16, false, false)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(result)
}

func TestOnlyWithDigits(t *testing.T) {
	result, err := GeneratePassword(16, true, false)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(result)
}

func TestOnlyWithDigitsAndSymbols(t *testing.T) {
	result, err := GeneratePassword(16, true, true)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(result)
}
