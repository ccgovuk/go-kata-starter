package main

import (
	"fmt"
)

type sku struct {
	quantity int
	discount int
	price    int
}

func main() {
	price := Scan("ABA", getDiscountsActual())
	fmt.Printf("%d\n", price)

}
func Scan(items string, skus map[rune]sku) int {

	discount := calculateDiscount(items, skus)
	total := 0
	for _, item := range items {
		total += skus[item].price
	}
	return total - discount

}

func getDiscountsActual() map[rune]sku {
	var discounts = make(map[rune]sku)
	discounts['A'] = sku{quantity: 3, discount: 20, price: 50}
	discounts['B'] = sku{quantity: 2, discount: 15, price: 30}
	discounts['C'] = sku{price: 20}
	discounts['D'] = sku{price: 15}
	return discounts
}

func calculateDiscount(items string, skus map[rune]sku) int {

	counts := getCounts(items)

	discount := 0

	discount = (counts['A'] / skus['A'].quantity) * skus['A'].discount
	discount += (counts['B'] / skus['B'].quantity) * skus['B'].discount

	return discount
}
func getCounts(items string) map[rune]int {
	var counts = make(map[rune]int)
	for _, item := range items {
		counts[item]++
	}
	return counts
}
