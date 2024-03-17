package service

import (
	"fmt"

	"payments.xws.com/model"
	"payments.xws.com/repo"
)

type TourPurchaseTokenService struct {
	TokenRepository *repo.TourPurchaseTokenRepository
}

func (service *TourPurchaseTokenService) GetAll() ([]model.TourPurchaseToken, error) {
	tokens, err := service.TokenRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tokens: %w", err)
	}

	return tokens, err
}

func (service *TourPurchaseTokenService) GetAllByTourist(touristId int) ([]model.TourPurchaseToken, error) {
	tokens, err := service.TokenRepository.GetAllByTourist(touristId)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tokens for tourist: %w", err)
	}

	return tokens, err
}
