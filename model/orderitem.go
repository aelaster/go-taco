package model

type OrderItem struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type OrderItems []OrderItem
