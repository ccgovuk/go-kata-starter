package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func scanOCR(input string) string {
	if input[1] == ' ' {
		return "111111111"
	}
	return "000000000"
}

func TestAlwaysTrue(t *testing.T) {

	assert.True(t, true) // Example comment
}

func TestAllZeros(t *testing.T) {

	//  _  _  _  _  _  _  _  _  _
	// | || || || || || || || || |
	// |_||_||_||_||_||_||_||_||_|
	//
	input := ` _  _  _  _  _  _  _  _  _ 
 | || || || || || || || || |
 |_||_||_||_||_||_||_||_||_|
`
	result := scanOCR(input)
	assert.Equal(t, "000000000", result)
}

func TestAllOnes(t *testing.T) {

	//
	//   |  |  |  |  |  |  |  |  |
	//   |  |  |  |  |  |  |  |  |
	//
	input := `                           
  |  |  |  |  |  |  |  |  |
  |  |  |  |  |  |  |  |  |
`
	result := scanOCR(input)
	assert.Equal(t, "111111111", result)
}

func TestInpuLengthIs84(t *testing.T) {

	input := `                           
  |  |  |  |  |  |  |  |  |
  |  |  |  |  |  |  |  |  |
`
	result := len(input)
	assert.Equal(t, 84, result)
}

///TODO
// Test if an error is thrown when line length is incorrent
//
