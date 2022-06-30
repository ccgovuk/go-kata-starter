package order

type Order struct {
	Total      int
	OrderItems []OrderItem
}

type OrderItem struct {
	Item      string
	PricePaid int
}
