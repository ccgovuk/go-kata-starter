package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlwaysTrue(t *testing.T) {

	assert.True(t, true)
}

func generateExpectedOutput() [100]string {
	var numbers [100]string
	for i := 1; i <= 100; i++ {
		numbers[i-1] = strconv.Itoa(i)
	}

	return numbers
}

func TestDivisibleByThreeReturnsFizz(t *testing.T) {
	positions := [...]int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48, 51, 54, 57, 60, 63, 66, 69, 72, 75, 78, 81, 84, 87, 90, 93, 96, 99}
	for i := 0; i < len(positions); i++ {
		assert.Equal(t, "fizz", fizzBuzz()[positions[i]-1])
	}

}

func TestDivisibleByFiveContainsBuzz(t *testing.T) {
	positions := [...]int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100}
	for i := 0; i < len(positions); i++ {
		assert.True(t, string.Contains(fizzBuzz()[positions[i]-1], "buzz"))
	}

}

func fizzBuzz() [100]string {
	var numbers [100]string

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			numbers[i-1] = "fizz"
		} else {
			numbers[i-1] = strconv.Itoa(i)
		}

	}
	return numbers
}
