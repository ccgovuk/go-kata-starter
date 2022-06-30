package receiptPrinter

import (
	"testing"

	"github.com/contino/go-kata-starter/order"
	"github.com/stretchr/testify/assert"
)

func TestPrintTotal(t *testing.T) {
	order := order.Order{Total: 50}
	assert.Contains(t, PrintReceipt(order), "50")
}

func TestPrintTotalFormatted(t *testing.T) {
	order := order.Order{Total: 50}
	expectedReceipt := "Total         £50.00\n"
	receipt := PrintReceipt(order)
	assert.Equal(t, expectedReceipt, receipt)
}
