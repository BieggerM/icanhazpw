package service

import (
	"fmt"
	"testing"
)

func TestPasswordLength(t *testing.T) {
	/* testing length 0 */
	result, err := GeneratePassword(0, false, false)
	if err == nil {
		fmt.Print("Expected error for input length 0")
		return
	}

	/* testing length */
	testCases := []int{6, 8, 2, 1, 88, 32, 1}
	for x, i := range testCases {
		if result, err = GeneratePassword(i, false, false); err != nil {
			fmt.Print("Error while evaluating")
		}
		if len(result) != i {
			fmt.Printf("Expected length %v for element in position %v, got %v", i, x, len(result))
		}
	}
}

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

func TestOnlyWithSymbols(t *testing.T) {
	result, err := GeneratePassword(16, false, true)
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
