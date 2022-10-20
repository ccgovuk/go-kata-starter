package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var hashMap = map[string]string{
	" _ | ||_|": "0",
	"     |  |": "1",
	" _  _||__": "2",
	" _  _| _|": "3",
	"   |_|  |": "4",
	" _ |_||_|": "8",
}

func scanOCR(input string) string {
	output := ""

	if len(input) != 84 {
		return " ERR"
	}

	// switch {
	// case input[1] == ' ':
	// 	return "111111111"
	// case input[29] == '_':
	// 	return "222222222"
	// default:
	// 	return "000000000"
	// }

	// for i, character := range inputArray[0] {
	// 	o := input[0:3]
	// 	o := input[3:6]
	// }

	for i := 0; i < 27; i += 3 {
		str1 := input[i : i+3]
		str2 := input[i+28 : i+31]
		str3 := input[i+56 : i+59]
		digit := parseDigit([]string{str1, str2, str3})
		output += digit
	}

	return output
}

func parseDigit(input []string) string {
	key := strings.Join(input, "")
	return hashMap[key]
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
	assert.Equal(t, "111111111", result)
}

func TestAllTwos(t *testing.T) {
	input := " _  _  _  _  _  _  _  _  _ \n"
	input += " _| _| _| _| _| _| _| _| _|\n"
	input += "|_ |_ |_ |_ |_ |_ |_ |_ |_ \n"

	result := scanOCR(input)
	assert.Equal(t, "222222222", result)
}

func TestIncorrectLengthReturnsError(t *testing.T) {

	input := "_  _  _  _  _  _  _  _  _ \n"
	input += "_| _| _| _| _| _| _| _| _|\n"
	input += "|_ |_ |_ |_ |_ |_ |_ |_ |_ \n"

	result := scanOCR(input)
	assert.Equal(t, " ERR", result[len(result)-4:])
}

func TestDigitEightReturnedFromComponantStrings(t *testing.T) {
	input := []string{" _ ", "|_|", "|_|"}
	result := parseDigit(input)
	assert.Equal(t, "8", result)
}

func TestDigitFourReturnedFromComponantStrings(t *testing.T) {
	input := []string{"   ", "|_|", "  |"}
	result := parseDigit(input)
	assert.Equal(t, "4", result)
}

// TODO: Parse digits
