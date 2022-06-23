package receiptPrinter

import (
	"fmt"

	"github.com/contino/go-kata-starter/basket"
)

func PrintReceipt(b basket.Basket) string {
	return fmt.Sprintf("Total         £%d.00\n", b.Total)
}
