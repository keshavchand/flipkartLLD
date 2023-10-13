package models

type Order struct {
	RestaurantName string
	ItemName       string
	Quantity       uint64
	Price          float32
}
