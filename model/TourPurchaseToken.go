package model

type TourPurchaseToken struct {
	ID        int `json:"id" gorm:"primaryKey"`
	TouristId int `json:"touristId"`
	IdTour    int `json:"idTour"`
}

func NewTourPurchaseToken(touristId, idTour int) *TourPurchaseToken {
	return &TourPurchaseToken{
		TouristId: touristId,
		IdTour:    idTour,
	}
}
