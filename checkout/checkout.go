package checkout

type Sku struct {
	Quantity int
	Discount int
	Price    int
}

func Scan(items string, skus map[rune]Sku) int {

	discount := calculateDiscount(items, skus)
	total := 0
	for _, item := range items {
		total += skus[item].Price
	}
	return total - discount

}

func calculateDiscount(items string, skus map[rune]Sku) int {

	counts := getCounts(items)

	discount := 0

	discount = (counts['A'] / skus['A'].Quantity) * skus['A'].Discount
	discount += (counts['B'] / skus['B'].Quantity) * skus['B'].Discount

	return discount
}
func getCounts(items string) map[rune]int {
	var counts = make(map[rune]int)
	for _, item := range items {
		counts[item]++
	}
	return counts
}
