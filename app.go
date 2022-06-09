package main

import (
	"os"

	"github.com/contino/go-kata-starter/checkout"
	"github.com/contino/go-kata-starter/receiptPrinter"
)

func main() {
	args := os.Args
	price := checkout.Scan(args[1], GetDiscountsActual())
	receiptPrinter.PrintReceipt(price)

}

func GetDiscountsActual() map[rune]checkout.Sku {
	var discounts = make(map[rune]checkout.Sku)
	discounts['A'] = checkout.Sku{Quantity: 3, Discount: 20, Price: 50}
	discounts['B'] = checkout.Sku{Quantity: 2, Discount: 15, Price: 30}
	discounts['C'] = checkout.Sku{Price: 20}
	discounts['D'] = checkout.Sku{Price: 15}
	return discounts
}
