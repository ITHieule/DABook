package types

type OrderStat struct {
	OrderDate      string  `json:"order_date"`
	TotalBooksSold int     `json:"total_books_sold"` // Tổng số sách đã bán
	TotalRevenue   float32 `json:"total_revenue"`
}
