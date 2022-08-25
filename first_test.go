package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func chop(target int, search_array []int) int {

	for index, num := range search_array {
		if num == target {
			return index
		}
	}
	return -1
}

func TestAlwaysTrue(t *testing.T) {
	assert.True(t, true)
}

func TestFindOneInListWithOnlyOne(t *testing.T) {
	ans := chop(1, []int{1})
	assert.Equal(t, 0, ans)
}

func TestFindThreeInIndexOne(t *testing.T) {
	ans := chop(3, []int{1, 3, 5})
	assert.Equal(t, 1, ans)
}

func TestFindTwoInIndexMinusOne(t *testing.T) {
	ans := chop(2, []int{1, 3, 5})
	assert.Equal(t, -1, ans)
}

func TestNiceLongArray(t *testing.T) {
	ans := chop(5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	assert.Equal(t, 4, ans)
}
