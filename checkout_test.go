package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlwaysTrue(t *testing.T) {

	assert.True(t, true)
}

func TestEmptyBasketIsZero(t *testing.T) {

	var TotalPrice = ScanWithPriceAndDiscount("")
	assert.Equal(t, 0, TotalPrice)
}

func TestItemAisFifty(t *testing.T) {

	var TotalPrice = ScanWithPriceAndDiscount("A")
	assert.Equal(t, 50, TotalPrice)
}

func TestItemAAisOneHundred(t *testing.T) {

	var TotalPrice = ScanWithPriceAndDiscount("AA")
	assert.Equal(t, 100, TotalPrice)
}

func TestItemBisThirty(t *testing.T) {

	var TotalPrice = ScanWithPriceAndDiscount("B")
	assert.Equal(t, 30, TotalPrice)

}

func TestItemAAAisOneThirty(t *testing.T) {
	var TotalPrice = ScanWithPriceAndDiscount("AAA")
	assert.Equal(t, 130, TotalPrice)
}

func TestItemAAAAisOneEighty(t *testing.T) {
	var TotalPrice = ScanWithPriceAndDiscount("AAAA")
	assert.Equal(t, 180, TotalPrice)
}

func TestItemAAAAAAis260(t *testing.T) {
	var TotalPrice = ScanWithPriceAndDiscount("AAAAAA")
	assert.Equal(t, 260, TotalPrice)
}

func TestItemBBis45(t *testing.T) {
	var TotalPrice = ScanWithPriceAndDiscount("BB")
	assert.Equal(t, 45, TotalPrice)
}

func TestItemABisEighty(t *testing.T) {

	var TotalPrice = ScanWithPriceAndDiscount("AB")
	assert.Equal(t, 80, TotalPrice)
}
func TestItemCBDAis115(t *testing.T) {

	var TotalPrice = ScanWithPriceAndDiscount("CBDA")
	assert.Equal(t, 115, TotalPrice)
}
func ScanWithPriceAndDiscount(items string) int {
	return Scan(items, getDiscounts())
}

func getDiscounts() map[rune]sku {
	var discounts = make(map[rune]sku)
	discounts['A'] = sku{quantity: 3, discount: 20, price: 50}
	discounts['B'] = sku{quantity: 2, discount: 15, price: 30}
	discounts['C'] = sku{price: 20}
	discounts['D'] = sku{price: 15}
	return discounts
}
