package repo

import (
	"gorm.io/gorm"
	"payments.xws.com/model"
)

// TourPurchaseTokenRepository handles database operations for TourPurchaseToken entities.
type TourPurchaseTokenRepository struct {
	DatabaseConnection *gorm.DB
}

// GetAll retrieves all TourPurchaseToken records from the database.
func (r *TourPurchaseTokenRepository) GetAll() ([]model.TourPurchaseToken, error) {
	var tokens []model.TourPurchaseToken
	result := r.DatabaseConnection.Find(&tokens)
	return tokens, result.Error
}

// HasToken checks if a specific TourPurchaseToken exists for a given touristId and tourId.
func (r *TourPurchaseTokenRepository) HasToken(touristId, tourId int) (bool, error) {
	var count int64
	result := r.DatabaseConnection.Model(&model.TourPurchaseToken{}).Where("tourist_id = ? AND id_tour = ?", touristId, tourId).Count(&count)
	return count > 0, result.Error
}

func (r *TourPurchaseTokenRepository) GetAllByTourist(touristId int) ([]model.TourPurchaseToken, error) {
	var tokens []model.TourPurchaseToken
	// Apply the filter before executing the Find method
	result := r.DatabaseConnection.Where("tourist_id = ?", touristId).Find(&tokens)
	return tokens, result.Error
}

func (r *TourPurchaseTokenRepository) Create(token *model.TourPurchaseToken) (model.TourPurchaseToken, error) {
	dbCreationResult := r.DatabaseConnection.Create(token)
	if dbCreationResult != nil {
		err := dbCreationResult.Error
		return *token, err // return *cart and the error if creation failed
	}

	return *token, nil
}
