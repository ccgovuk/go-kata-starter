package main

import (
	"strings"
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

// func TestAllOnes(t *testing.T) {

// 	input := "                           \n"
// 	input += "  |  |  |  |  |  |  |  |  |\n"
// 	input += "  |  |  |  |  |  |  |  |  |\n"

// 	result := scanOCR(input)
// 	assert.Equal(t, "111111111 ERR", result)
// }

// func TestAllTwos(t *testing.T) {
// 	input := " _  _  _  _  _  _  _  _  _ \n"
// 	input += " _| _| _| _| _| _| _| _| _|\n"
// 	input += "|_ |_ |_ |_ |_ |_ |_ |_ |_ \n"

// 	result := scanOCR(input)
// 	assert.Equal(t, "222222222 ERR", result)
// }

// func TestIncorrectLengthReturnsError(t *testing.T) {

// 	input := "_  _  _  _  _  _  _  _  _ \n"
// 	input += "_| _| _| _| _| _| _| _| _|\n"
// 	input += "|_ |_ |_ |_ |_ |_ |_ |_ |_ \n"

// 	result := scanOCR(input)
// 	assert.Equal(t, " ERR", result[len(result)-4:])
// }

func TestDigitEightReturnedFromComponentStrings(t *testing.T) {
	input := []string{" _ ", "|_|", "|_|"}
	digit := Digit{asciiChar: input}
	digit.parseDigit()
	assert.Equal(t, "8", digit.strValue)
}

func TestDigitFourReturnedFromComponentStrings(t *testing.T) {
	input := []string{"   ", "|_|", "  |"}
	digit := Digit{asciiChar: input}
	digit.parseDigit()
	assert.Equal(t, "4", digit.strValue)
}

func TestAllDigitsParsedCorrectly(t *testing.T) {
	input := "    _  _     _  _  _  _  _ \n"
	input += "  | _| _||_||_ |_   ||_||_|\n"
	input += "  ||_  _|  | _||_|  ||_| _|\n"

	result := scanOCR(input)
	assert.Equal(t, "123456789", result)
}

func TestCheckSumPass(t *testing.T) {
	input := [9]Digit{
		{intValue: 4},
		{intValue: 5},
		{intValue: 7},
		{intValue: 5},
		{intValue: 0},
		{intValue: 8},
		{intValue: 0},
		{intValue: 0},
		{intValue: 0},
	}
	accountNumber := AccountNumber{digits: input}
	checksum := accountNumber.calculateCheckSum()
	assert.True(t, checksum)
}

// func TestCheckSumFail(t *testing.T) {
// 	input := []int{6, 6, 4, 3, 7, 1, 4, 9, 5}
// 	checksum := calculateCheckSum(input)
// 	assert.False(t, checksum)
// }

// func TestConvertStringToDigits(t *testing.T) {
// 	input := "457508000"
// 	result, _ := parseAccountStringToDigits(input)
// 	assert.Equal(t, []int{4, 5, 7, 5, 0, 8, 0, 0, 0}, result)
// }

func TestDigitsParsedCheckSumCorrect(t *testing.T) {
	input := "    _  _     _  _  _  _  _ \n"
	input += "  | _| _||_||_ |_   ||_||_|\n"
	input += "  ||_  _|  | _||_|  ||_| _|\n"

	result := scanOCR(input)
	assert.Equal(t, "123456789", result)
}

// func TestDigitsParsedCheckSumInCorrect(t *testing.T) {
// 	input := "    _  _     _  _  _  _  _ \n"
// 	input += "  | _| _||_||_ |_   ||_||_|\n"
// 	input += "  ||_  _|  | _||_|  ||_||_|\n"

// 	result := scanOCR(input)
// 	assert.Equal(t, "123456788 ERR", result)
// }

func TestDigitsParsedIllegalCharacters(t *testing.T) {
	input := "    _  _     _  _  _  _    \n"
	input += "  | _| _||_||_ |_   ||_||_|\n"
	input += "  ||_  _|  | _||_|  ||_||_|\n"

	result := scanOCR(input)
	assert.Equal(t, "12345678? ILL", result)
}

func TestParseIllegalCharacterInMiddle(t *testing.T) {
	input := "    _  _        _  _  _    \n"
	input += "  | _| _||_||_ |_   ||_||_|\n"
	input += "  ||_  _|  | _||_|  ||_||_|\n"

	result := scanOCR(input)
	assert.Equal(t, "1234?678? ILL", result)
}

func TestAllOnesVersion2(t *testing.T) {

	input := "                           \n"
	input += "  |  |  |  |  |  |  |  |  |\n"
	input += "  |  |  |  |  |  |  |  |  |\n"

	result := scanOCR(input)
	assert.Equal(t, "711111111", result)
}

func TestAllOnesVersion3(t *testing.T) {

	input := " _  _  _  _  _  _  _  _  _ \n"
	input += "  |  |  |  |  |  |  |  |  |\n"
	input += "  |  |  |  |  |  |  |  |  |\n"

	result := scanOCR(input)
	assert.Equal(t, "777777177", result)
}

func TestInvalidDigit(t *testing.T) {
	input := "|_ "
	input += "| |"
	input += "|_|"

	digit := MakeDigit([]string{input})
	// digit := MakeDigit(strings.Split(input, ""))

	result := digit.isValid()

	assert.Equal(t, false, result)
}

func TestValidDigit(t *testing.T) {
	input := " _ "
	input += "| |"
	input += "|_|"

	digit := MakeDigit(strings.Split(input, ""))

	result := digit.isValid()

	assert.Equal(t, true, result)
}

func TestAlternates(t *testing.T) {
	input := " _ "
	input += "|_|"
	input += "|_|"

	digit := MakeDigit([]string{input})
	result := digit.alternates()

	assert.EqualValues(t, []string{"0", "6", "9"}, result)
}

func TestAlternatesAgain(t *testing.T) {
	input := " _ "
	input += "| |"
	input += "|_|"

	digit := MakeDigit([]string{input})
	result := digit.alternates()

	assert.EqualValues(t, []string{"8"}, result)
}

func TestDigitToAsciiOk(t *testing.T) {
	input := "0"
	expected := " _ | ||_|"

	assert.Equal(t, expected, digitToAscii(input))
}

func TestDigitToAsciiFail(t *testing.T) {
	input := "a"
	expected := ""

	assert.Equal(t, expected, digitToAscii(input))
}
