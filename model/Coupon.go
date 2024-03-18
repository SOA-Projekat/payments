package model

import (
	"math/rand"
	"time"
)

type Coupon struct {
	ID             int       `json:"id" gorm:"primaryKey"`
	Code           string    `json:"code"`
	Discount       int       `json:"discount"`
	ExpirationDate time.Time `json:"expirationDate"`
	TourId         int       `json:"tourId"`
	TouristId      int       `json:"touristId"`
	AuthorId       int       `json:"authorId"`
}

func NewCoupon(code string, discount int, expirationDate time.Time, tourId, touristId, authorId int) *Coupon {
	coupon := &Coupon{
		Code:           code,
		Discount:       discount,
		ExpirationDate: expirationDate,
		TourId:         tourId,
		TouristId:      touristId,
		AuthorId:       authorId,
	}
	return coupon
}

func (c *Coupon) IsExpired() bool {
	return c.ExpirationDate.Before(time.Now())
}

func (c *Coupon) GenerateCode() {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomChars := make([]byte, 8)
	for i := range randomChars {
		randomChars[i] = charset[seededRand.Intn(len(charset))]
	}
	c.Code = string(randomChars)
}
