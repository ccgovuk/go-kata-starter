package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlwaysTrue(t *testing.T) {

	assert.True(t, true) // Example comment
}

func TestAllZeros(t *testing.T) {

	input := " _  _  _  _  _  _  _  _  _ \n"
	input += "| || || || || || || || || |\n"
	input += "|_||_||_||_||_||_||_||_||_|\n"

	result := scanOCR(input)
	assert.Equal(t, "000000000", result)
}

func TestAllOnes(t *testing.T) {

	input := "                           \n"
	input += "  |  |  |  |  |  |  |  |  |\n"
	input += "  |  |  |  |  |  |  |  |  |\n"

	result := scanOCR(input)
	assert.Equal(t, "111111111 ERR", result)
}

func TestAllTwos(t *testing.T) {
	input := " _  _  _  _  _  _  _  _  _ \n"
	input += " _| _| _| _| _| _| _| _| _|\n"
	input += "|_ |_ |_ |_ |_ |_ |_ |_ |_ \n"

	result := scanOCR(input)
	assert.Equal(t, "222222222 ERR", result)
}

func TestIncorrectLengthReturnsError(t *testing.T) {

	input := "_  _  _  _  _  _  _  _  _ \n"
	input += "_| _| _| _| _| _| _| _| _|\n"
	input += "|_ |_ |_ |_ |_ |_ |_ |_ |_ \n"

	result := scanOCR(input)
	assert.Equal(t, " ERR", result)
}

func TestDigitEightReturnedFromComponentStrings(t *testing.T) {
	input := []string{" _ ", "|_|", "|_|"}
	result := parseDigit(input)
	assert.Equal(t, "8", result)
}

func TestDigitFourReturnedFromComponentStrings(t *testing.T) {
	input := []string{"   ", "|_|", "  |"}
	result := parseDigit(input)
	assert.Equal(t, "4", result)
}

func TestAllDigitsParsedCorrectly(t *testing.T) {
	input := "    _  _     _  _  _  _  _ \n"
	input += "  | _| _||_||_ |_   ||_||_|\n"
	input += "  ||_  _|  | _||_|  ||_| _|\n"

	result := scanOCR(input)
	assert.Equal(t, "123456789", result)
}

func TestCheckSumPass(t *testing.T) {
	input := []int{4, 5, 7, 5, 0, 8, 0, 0, 0}
	checksum := calculateCheckSum(input)
	assert.True(t, checksum)
}

func TestCheckSumFail(t *testing.T) {
	input := []int{6, 6, 4, 3, 7, 1, 4, 9, 5}
	checksum := calculateCheckSum(input)
	assert.False(t, checksum)
}

func TestConvertStringToDigits(t *testing.T) {
	input := "457508000"
	result, _ := parseAccountStringToDigits(input)
	assert.Equal(t, []int{4, 5, 7, 5, 0, 8, 0, 0, 0}, result)
}

func TestDigitsParsedCheckSumCorrect(t *testing.T) {
	input := "    _  _     _  _  _  _  _ \n"
	input += "  | _| _||_||_ |_   ||_||_|\n"
	input += "  ||_  _|  | _||_|  ||_| _|\n"

	result := scanOCR(input)
	assert.Equal(t, "123456789", result)
}

func TestDigitsParsedCheckSumInCorrect(t *testing.T) {
	input := "    _  _     _  _  _  _  _ \n"
	input += "  | _| _||_||_ |_   ||_||_|\n"
	input += "  ||_  _|  | _||_|  ||_||_|\n"

	result := scanOCR(input)
	assert.Equal(t, "123456788 ERR", result)
}

func TestDigitsParsedIllegalCharacters(t *testing.T) {
	input := "    _  _     _  _  _  _    \n"
	input += "  | _| _||_||_ |_   ||_||_|\n"
	input += "  ||_  _|  | _||_|  ||_||_|\n"

	result := scanOCR(input)
	assert.Equal(t, "12345678? ILL", result)
}

func TestDigitsParsedIllegalCharactersMiddle(t *testing.T) {
	input := "    _  _     _  _  _  _    \n"
	input += "  | _| _|| ||_ |_   ||_||_|\n"
	input += "  ||_  _|  | _||_|  ||_||_|\n"

	result := scanOCR(input)
	assert.Equal(t, "123?5678? ILL", result)
}
