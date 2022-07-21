package order

type Order struct {
	Total      int
	OrderItems []OrderItem
}

type OrderItem struct {
	Item      rune
	PricePaid int
}
