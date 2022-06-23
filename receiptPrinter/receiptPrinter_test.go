package receiptPrinter

import (
	"testing"

	"github.com/contino/go-kata-starter/basket"
	"github.com/stretchr/testify/assert"
)

func TestPrintTotal(t *testing.T) {
	basket := basket.Basket{Total: 50}
	assert.Contains(t, PrintReceipt(basket), "50")
}

func TestPrintTotalFormatted(t *testing.T) {
	basket := basket.Basket{Total: 50}
	expectedReceipt := "Total         £50.00\n"
	receipt := PrintReceipt(basket)
	assert.Equal(t, expectedReceipt, receipt)
}
