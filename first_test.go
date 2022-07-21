package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlwaysTrue(t *testing.T) {

	assert.True(t, true)
}

func TestKinshipPinkish(t *testing.T) {
	result := solve("kinship")

	assert.Contains(t, result, "pinkish")
}

func TestPoolLoopPolo(t *testing.T) {
	result := solve("loop")

	assert.Equal(t, len(result), 2)
	assert.Contains(t, result, "polo")
	assert.Contains(t, result, "pool")
}

func solve(word string) []string {
	dictionary := []string{"pinkish", "polo", "pool"}
	result := []string{}
	x := runeArray(word)

	for _, wordInDictonary := range dictionary {
		if runeArray(wordInDictonary) == x {
			result = append(result, wordInDictonary)
		}
	}
	return result
}

func runeArray(word string) string {
	runes := []rune(word)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
