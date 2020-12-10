package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlwaysTrue(t *testing.T) {

	assert.True(t, true)
}

func TestShouldReturn1for1(t *testing.T) {
	result := fizzBuzz()

	assert.Equal(t, result[0], "1")
}

func TestShouldReturnFizzfor2ndElement(t *testing.T) {
	result := fizzBuzz()

	assert.Equal(t, result[2], "Fizz")
}

func TestShouldReturnBuzzfor4thElement(t *testing.T) {
	result := fizzBuzz()

	assert.Equal(t, result[4], "Buzz")
}

func TestShouldReturnBuzzfor5thElement(t *testing.T) {
	result := fizzBuzz()

	assert.Equal(t, result[5], "Fizz")
}

func TestShouldReturnExpectedOutput(t *testing.T) {
	result := fizzBuzz()
	assert.Equal(t, result, [10]string{"1",
		"2",
		"Fizz",
		"4",
		"Buzz",
		"Fizz",
		"7",
		"8",
		"Fizz",
		"Buzz"})
}

func fizzBuzz() [10]string {

	results := [10]string{}
	for i := 0; i < 10; i++ {
		results[i] = strconv.Itoa(i + 1)
		if (i+1)%3 == 0 {
			results[i] = "Fizz"
		} else if (i+1)%5 == 0 {
			results[i] = "Buzz"
		} else {

		}
	}
	return results
}
