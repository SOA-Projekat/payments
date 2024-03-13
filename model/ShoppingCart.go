package model

type ShoppingCart struct {
	ID         int         `json:"id" gorm:"primaryKey"`
	TouristId  int         `json:"touristId"`
	OrderItems []OrderItem `json:"orderItems" gorm:"type:jsonb;default:'[]'"`
	Total      float64     `json:"total"`
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
