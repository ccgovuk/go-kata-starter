package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlwaysTrue(t *testing.T) {

	assert.True(t, true)
}

func TestShouldReturn1for1(t *testing.T) {
	result := FizzBuzz(10)

	assert.Equal(t, "1",result[0])
}

func TestShouldReturnFizzfor2ndElement(t *testing.T) {
	result := FizzBuzz(10)

	assert.Equal(t,"Fizz", result[2] )
}

func TestShouldReturnBuzzfor4thElement(t *testing.T) {
	result := FizzBuzz(10)

	assert.Equal(t, "Buzz",result[4])
}

func TestShouldReturnBuzzfor5thElement(t *testing.T) {
	result := FizzBuzz(10)

	assert.Equal(t,"Fizz", result[5])
}

func TestShouldReturnFizzBuzzfor15thElement(t *testing.T){
	result := FizzBuzz(15)

	assert.Equal(t, "FizzBuzz", result[14])
}
func TestShouldReturnExpectedOutput(t *testing.T) {
	result := FizzBuzz(10)
	assert.Equal(t,[]string{"1",
		"2",
		"Fizz",
		"4",
		"Buzz",
		"Fizz",
		"7",
		"8",
		"Fizz",
		"Buzz"}, result )
}

