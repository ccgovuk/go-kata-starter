package checkout

import "github.com/contino/go-kata-starter/order"

type Sku struct {
	Quantity int
	Discount int
	Price    int
}

func Scan(itemsInBasket string, skus map[rune]Sku) order.Order {

	discount := calculateDiscount(itemsInBasket, skus)
	total := 0
	scannedOrder := order.Order{}
	for _, itemInBasket := range itemsInBasket {
		total += skus[itemInBasket].Price
		scannedOrder.OrderItems = append(scannedOrder.OrderItems, order.OrderItem{Item: itemInBasket, PricePaid: skus[itemInBasket].Price})

	}
	scannedOrder.Total = total - discount

	return scannedOrder
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
