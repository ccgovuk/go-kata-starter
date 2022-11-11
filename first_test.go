package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func scanOCR(input string) string {
	output := ""

	if len(input) != 84 {
		return " ERR"
	}

	for i := 0; i < 27; i += 3 {
		str1 := input[i : i+3]
		str2 := input[i+28 : i+31]
		str3 := input[i+56 : i+59]
		digit := parseDigit([]string{str1, str2, str3})
		output += digit
	}
	digits, err := parseAccountStringToDigits(output)
	if err != nil {
		output += " ILL"
		return output
	}
	if !calculateCheckSum(digits) {
		output += " ERR"
	}
	return output
}

func parseDigit(input []string) string {
	key := strings.Join(input, "")
	value, ok := hashMap[key]
	if ok {
		return value
	}
	return "?"
}

func calculateCheckSum(input []int) bool {
	sum := 0
	for i, j := range input {
		sum += j * (len(input) - i)
	}

	return sum%11 == 0
}

func parseAccountStringToDigits(input string) ([]int, error) {
	arrayOfStrings := strings.Split(input, "")
	arrayOfInts := make([]int, len(arrayOfStrings))

	for i, s := range arrayOfStrings {
		num, err := strconv.Atoi(s)
		if err != nil {
			return []int{}, err
		}

		arrayOfInts[i] = num
	}

	return arrayOfInts, nil
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

func TestParseIllegalCharacterInMiddle(t *testing.T) {
	input := "    _  _        _  _  _    \n"
	input += "  | _| _||_||_ |_   ||_||_|\n"
	input += "  ||_  _|  | _||_|  ||_||_|\n"

	result := scanOCR(input)
	assert.Equal(t, "1234?678? ILL", result)
}
