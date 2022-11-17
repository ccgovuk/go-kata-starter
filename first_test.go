package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AccountNumber struct {
	digits [9]Digit
}
type Digit struct {
	value     string
	intValue  int
	asciiChar []string
}

var hashMap = map[string]string{
	" _ | ||_|": "0",
	"     |  |": "1",
	" _  _||_ ": "2",
	" _  _| _|": "3",
	"   |_|  |": "4",
	" _ |_  _|": "5",
	" _ |_ |_|": "6",
	" _   |  |": "7",
	" _ |_||_|": "8",
	" _ |_| _|": "9",
}

func (d *Digit) parseDigit() {
	key := strings.Join(d.asciiChar, "")
	value, ok := hashMap[key]
	if ok {
		d.value = value
	} else {
		d.value = "?"
	}
}

func scanOCR(input string) string {
	output := ""

	if len(input) != 84 {
		return " ERR"
	}

	accountNumber := AccountNumber{}

	for i := 0; i < 27; i += 3 {
		str1 := input[i : i+3]
		str2 := input[i+28 : i+31]
		str3 := input[i+56 : i+59]
		digit := Digit{asciiChar: []string{str1, str2, str3}}
		digit.parseDigit()
		accountNumber.digits[i] = digit
		output += digit.value
	}
	err := accountNumber.parseAccountStringToDigits()
	if err != nil {
		output += " ILL"
		return output
	}
	if !accountNumber.calculateCheckSum() {
		output += " ERR"
	}
	return output
}

func (ocr *AccountNumber) calculateCheckSum() bool {
	sum := 0
	for i, j := range ocr.digits {
		sum += j.intValue * (len(ocr.digits) - i)
	}

	return sum%11 == 0
}

func (ocr *AccountNumber) parseAccountStringToDigits() error {

	for i, s := range ocr.digits {
		num, err := strconv.Atoi(s.value)
		if err != nil {
			return err
		}

		ocr.digits[i].intValue = num
	}

	return nil
}

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
	assert.Equal(t, " ERR", result[len(result)-4:])
}

func TestDigitEightReturnedFromComponentStrings(t *testing.T) {
	input := []string{" _ ", "|_|", "|_|"}
	digit := Digit{asciiChar: input}
	digit.parseDigit()
	assert.Equal(t, "8", digit.value)
}

func TestDigitFourReturnedFromComponentStrings(t *testing.T) {
	input := []string{"   ", "|_|", "  |"}
	digit := Digit{asciiChar: input}
	digit.parseDigit()
	assert.Equal(t, "4", digit.value)
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
		Digit{intValue: 4},
		Digit{intValue: 5},
		Digit{intValue: 7},
		Digit{intValue: 5},
		Digit{intValue: 0},
		Digit{intValue: 8},
		Digit{intValue: 0},
		Digit{intValue: 0},
		Digit{intValue: 0},
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

func TestParseIllegalCharacterInMiddle(t *testing.T) {
	input := "    _  _        _  _  _    \n"
	input += "  | _| _||_||_ |_   ||_||_|\n"
	input += "  ||_  _|  | _||_|  ||_||_|\n"

	result := scanOCR(input)
	assert.Equal(t, "1234?678? ILL", result)
}
