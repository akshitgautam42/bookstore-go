package models

type Book struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Quantity    uint   `json:"quantity"`
}

type OrderStatus string

const (
	OrderPlaced      OrderStatus = "PLACED"
	OrderProcessed   OrderStatus = "PROCESSED"
	OrderShipped     OrderStatus = "SHIPPED"
	OrderDelivered   OrderStatus = "DELIVERED"
	OrderCancelled   OrderStatus = "CANCELLED"
)

type Order struct {
	ID          uint        `json:"id"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	Email       string      `json:"email"`
	BookID      uint        `json:"book_id"`
	Book        Book        `json:"book"`
	Quantity    uint        `json:"quantity"`
	TotalAmount uint        `json:"total_amount"`
	Status      OrderStatus `json:"status"`
	CreatedAt   string      `json:"created_at"`
}
