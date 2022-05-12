package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetPrices() map[rune]int {
	var prices = make(map[rune]int)
	prices['A'] = 50
	prices['B'] = 30
	return prices
}

func TestAlwaysTrue(t *testing.T) {

	assert.True(t, true)
}

func TestEmptyBasketIsZero(t *testing.T) {

	var TotalPrice = Scan("", GetPrices())
	assert.Equal(t, 0, TotalPrice)
}

func TestItemAisFifty(t *testing.T) {

	var TotalPrice = Scan("A", GetPrices())
	assert.Equal(t, 50, TotalPrice)
}

func TestItemAAisOneHundred(t *testing.T) {

	var TotalPrice = Scan("AA", GetPrices())
	assert.Equal(t, 100, TotalPrice)
}

func TestItemBisThirty(t *testing.T) {

	var TotalPrice = Scan("B", GetPrices())
	assert.Equal(t, 30, TotalPrice)

}

func TestItemAAAisOneThirty(t *testing.T) {
	var TotalPrice = Scan("AAA", GetPrices())
	assert.Equal(t, 130, TotalPrice)
}

func TestItemAAAAisOneEighty(t *testing.T) {
	var TotalPrice = Scan("AAAA", GetPrices())
	assert.Equal(t, 180, TotalPrice)
}

func TestItemAAAAAAis260(t *testing.T) {
	var TotalPrice = Scan("AAAAAA", GetPrices())
	assert.Equal(t, 260, TotalPrice)
}

func TestItemABisEighty(t *testing.T) {

	var TotalPrice = Scan("AB", GetPrices())
	assert.Equal(t, 80, TotalPrice)
}

func Scan(items string, prices map[rune]int) int {
	discount := CalculateDiscount(items)
	total := 0
	for _, item := range items {
		total += prices[item]
	}
	return total - discount

}

func CalculateDiscount(items string) int {
	counts := GetCounts(items)

	discount := 0

	discount = (counts['A'] / 3) * 20

	return discount
}

func GetCounts(items string) map[rune]int {
	var counts = make(map[rune]int)
	for _, item := range items {
		counts[item]++
	}
	return counts
}
