package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlwaysTrue(t *testing.T) {

	assert.True(t, true)
}

func TestEmptyBasketIs0(t *testing.T) {

	var TotalPrice = scanWithPriceAndDiscount("")
	assert.Equal(t, 0, TotalPrice)
}

func TestItemAis50(t *testing.T) {

	var TotalPrice = scanWithPriceAndDiscount("A")
	assert.Equal(t, 50, TotalPrice)
}

func TestItemAAis100(t *testing.T) {

	var TotalPrice = scanWithPriceAndDiscount("AA")
	assert.Equal(t, 100, TotalPrice)
}

func TestItemBis30(t *testing.T) {

	var TotalPrice = scanWithPriceAndDiscount("B")
	assert.Equal(t, 30, TotalPrice)

}

func TestItemAAAis130(t *testing.T) {
	var TotalPrice = scanWithPriceAndDiscount("AAA")
	assert.Equal(t, 130, TotalPrice)
}

func TestItemAAAAis180(t *testing.T) {
	var TotalPrice = scanWithPriceAndDiscount("AAAA")
	assert.Equal(t, 180, TotalPrice)
}

func TestItemAAAAAAis260(t *testing.T) {
	var TotalPrice = scanWithPriceAndDiscount("AAAAAA")
	assert.Equal(t, 260, TotalPrice)
}

func TestItemBBis45(t *testing.T) {
	var TotalPrice = scanWithPriceAndDiscount("BB")
	assert.Equal(t, 45, TotalPrice)
}

func TestItemABis80(t *testing.T) {

	var TotalPrice = scanWithPriceAndDiscount("AB")
	assert.Equal(t, 80, TotalPrice)
}
func TestItemCBDAis115(t *testing.T) {

	var TotalPrice = scanWithPriceAndDiscount("CBDA")
	assert.Equal(t, 115, TotalPrice)
}
func scanWithPriceAndDiscount(items string) int {
	order := Scan(items, getDiscounts())
	return order.Total
}

func getDiscounts() map[rune]Sku {
	var discounts = make(map[rune]Sku)
	discounts['A'] = Sku{Quantity: 3, Discount: 20, Price: 50}
	discounts['B'] = Sku{Quantity: 2, Discount: 15, Price: 30}
	discounts['C'] = Sku{Price: 20}
	discounts['D'] = Sku{Price: 15}
	return discounts
}
