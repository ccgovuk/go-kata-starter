package main

import (
	"fmt"
	"strconv"
	"strings"
)

type AccountNumber struct {
	Digits []Digit
}

func (ac AccountNumber) String() string {
	output := ""
	for l := 0; l < 3; l++ {
		line := ""
		for _, d := range ac.Digits {
			line += d.AsciiLines[l]
		}
		output += fmt.Sprintf("%v\n", line)
	}
	return output
}

func (ac AccountNumber) Eval() string {
	output := ""
	for _, d := range ac.Digits {
		output += d.Eval()
	}
	return output
}


type Digit struct {
	AsciiLines []string
}

func (d Digit) String() string {
	output := ""
	for _, line := range d.AsciiLines {
		output += fmt.Sprintf("%v\n", line)
	}
	return output
}

func (d Digit) Eval() string {
	key := strings.Join(d.AsciiLines, "")
	value, ok := hashMap[key]
	if ok {
		return value
	}
	return "?"
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

func scanOCR(input string) string {
	output := ""
	accountNumber := AccountNumber{}

	if len(input) != 84 {
		return " ERR"
	}

	for i := 0; i < 27; i += 3 {
		str1 := input[i : i+3]
		str2 := input[i+28 : i+31]
		str3 := input[i+56 : i+59]
		d := Digit{AsciiLines: []string{str1, str2, str3}}
		accountNumber.Digits = append(accountNumber.Digits, d)
	}
	output = accountNumber.Eval()
	fmt.Println(accountNumber)
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
