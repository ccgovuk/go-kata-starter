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

func TestReceiptHasOne_A_WhenOnlyOne_A_InOrder(t *testing.T) {

	ois := []order.OrderItem{
		{Item: "A", PricePaid: 50},
	}

	o := order.Order{Total: 50, OrderItems: ois}

	expectedLine := "A               £50.00"
	receipt := PrintReceipt(o)
	assert.Contains(t, expectedLine, receipt)
}
