package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type ShoppingCart struct {
	ID         int             `json:"id" gorm:"primaryKey"`
	TouristId  int             `json:"touristId"`
	OrderItems OrderItemsValue `json:"orderItems" gorm:"type:jsonb;default:'[]'"`
	Total      float64         `json:"total"`
}

func NewShoppingCart(touristId int) *ShoppingCart {
	return &ShoppingCart{
		TouristId:  touristId,
		OrderItems: make([]OrderItem, 0),
		Total:      0,
	}
}

// AddOrderItem adds a new order item to the shopping cart.
func (sc *ShoppingCart) AddOrderItem(item OrderItem) bool {
	for _, existingItem := range sc.OrderItems {
		// Assuming equality is determined by the IDTour; adjust as necessary.
		if existingItem.IdTour == item.IdTour {
			return false // Item already exists
		}
	}
	sc.OrderItems = append(sc.OrderItems, item)
	sc.calculateTotal() // Recalculate the total
	return true
}

// RemoveOrderItem removes an order item from the shopping cart.
func (sc *ShoppingCart) RemoveOrderItem(idTour int) {
	for i, item := range sc.OrderItems {
		if item.IdTour == idTour {
			sc.OrderItems = append(sc.OrderItems[:i], sc.OrderItems[i+1:]...)
			sc.calculateTotal() // Recalculate the total
			return
		}
	}
}

// calculateTotal recalculates the total price of the shopping cart.
func (sc *ShoppingCart) calculateTotal() {
	total := 0.0
	for _, item := range sc.OrderItems {
		total += item.Price
	}
	sc.Total = total
}

type OrderItemsValue []OrderItem

// Value implements the driver.Valuer interface, converting the slice to JSON
func (o OrderItemsValue) Value() (driver.Value, error) {
	if len(o) == 0 {
		return "[]", nil // Return empty JSON array if slice is empty
	}
	return json.Marshal(o)
}

// Scan implements the sql.Scanner interface, converting JSON to a slice
func (o *OrderItemsValue) Scan(input interface{}) error {
	if input == nil {
		return nil // If the input is nil, just leave the slice as nil
	}

	bytes, ok := input.([]byte)
	if !ok {
		return fmt.Errorf("need []byte type for OrderItemsValue, got %T", input)
	}

	result := OrderItemsValue{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		return err
	}

	*o = result
	return nil
}
