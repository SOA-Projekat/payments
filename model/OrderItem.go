package model

type OrderItem struct {
	TourName string  `json:"tourName"`
	Price    float64 `json:"price"`
	IdTour   int     `json:"idTour"`
}
