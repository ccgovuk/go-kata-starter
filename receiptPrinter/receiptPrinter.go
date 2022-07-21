package receiptPrinter

import (
	"fmt"
	"strings"

	"github.com/contino/go-kata-starter/order"
)

func PrintReceipt(b order.Order) string {
	var reciept strings.Builder
	for _, item := range b.OrderItems {
		fmt.Fprintf(&reciept, "%c               £%d.00\n", item.Item, item.PricePaid)

	}
	fmt.Fprintf(&reciept, "Total         £%d.00\n", b.Total)
	return reciept.String()
}
