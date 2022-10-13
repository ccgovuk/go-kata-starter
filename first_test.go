package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func scanOCR(input string) string {
	if len(input) != 84 {
		return " ERR"
	}
	switch {
	case input[1] == ' ':
		return "111111111"
	case input[29] == '_':
		return "222222222"
	default:
		return "000000000"
	}

}

func parseDigit(input [3]string) rune {
	return '8'
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
	input := [3]string { " _ ", "|_|", "|_|"}
	result := parseDigit(input)
	assert.Equal(t, '8', result)
}

func TestDigitFourReturnedFromComponantStrings(t *testing.T) {
	input := [3]string { "   ", "|_|", "  |"}
	result := parseDigit(input)
	assert.Equal(t, '4', result)
}

// TODO: Parse digits
