package model

type CartItem struct {
	ID        int   `json:"id"`
	UserID    int64 `json:"user_id"`
	ProductID int   `json:"product_id"`
	Quantity  int   `json:"quantity"`
}
