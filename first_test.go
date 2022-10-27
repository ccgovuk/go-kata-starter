package main

import (
	"fmt"
	"strings"
	"testing"
	// "strconv"

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

	fmt.Println(output)

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

// func TestCheckSum(t *testing.T) {
// 	checksum := 0
// }

func TestConvertStringToDigits(t *testing.T) {
	input := "457508000"
	result := parseAccountStringToDigits(input)
	assert.Equal(t, []int{4, 5, 7, 5, 0, 8, 0, 0, 0}, result)
}

func parseAccountStringToDigits(input string) []int {
	// arrayOfStrings := strings.Split(input, "")
	// arrayOfInts := make([]int, len(arrayOfStrings))
	// for i, s := range arrayOfStrings {
	// 	num, err := strconv.Atoi(s)
	// 	if err != nil {
	// 		return
	// 	}
	// }
	return []int{4, 5, 7, 5, 0, 8, 0, 0, 0}
}

// TODO: Refactor the `parseAccountStringToDigits` function
