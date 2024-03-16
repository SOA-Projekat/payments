package repo

import (
	"gorm.io/gorm"
	"payments.xws.com/model"
)

type ShoppingCartRepo struct {
	DatabaseConnection *gorm.DB
}

func (repository *ShoppingCartRepo) GetByUserId(touristId int) (model.ShoppingCart, error) {
	cart := model.ShoppingCart{}
	databaseResult := repository.DatabaseConnection.First(&cart, "tourist_id = ?", touristId)
	if databaseResult != nil {
		return cart, databaseResult.Error
	}
	return cart, nil
}

func (repository *ShoppingCartRepo) Create(cart *model.ShoppingCart) (model.ShoppingCart, error) {
	dbCreationResult := repository.DatabaseConnection.Create(cart)
	if dbCreationResult != nil {
		err := dbCreationResult.Error
		return *cart, err // return *cart and the error if creation failed
	}

	return *cart, nil
}

func (repository *ShoppingCartRepo) Update(cart *model.ShoppingCart) (*model.ShoppingCart, error) {

	updateResult := repository.DatabaseConnection.Model(cart).Updates(cart)
	if updateResult != nil {
		return nil, updateResult.Error
	}

	return cart, nil
}

func (repository *ShoppingCartRepo) Get(id int) (model.ShoppingCart, error) {
	cart := model.ShoppingCart{}
	databaseResult := repository.DatabaseConnection.First(&cart, "id = ?", id)
	if databaseResult != nil {
		return cart, databaseResult.Error
	}
	return cart, nil
}
