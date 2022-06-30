package receiptPrinter

import (
	"fmt"

	"github.com/contino/go-kata-starter/order"
)

func PrintReceipt(b order.Order) string {
	return fmt.Sprintf("Total         £%d.00\n", b.Total)
}
