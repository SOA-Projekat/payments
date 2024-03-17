package service

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"payments.xws.com/model"
	"payments.xws.com/repo"
)

type ShoppingCartService struct {
	ShoppingCartRepo *repo.ShoppingCartRepo
	TokenRepo        *repo.TourPurchaseTokenRepository
}

func (service *ShoppingCartService) GetByUserId(touristId int) (model.ShoppingCart, error) {
	// Attempt to get the ShoppingCart by the user ID.
	cart, err := service.ShoppingCartRepo.GetByUserId(touristId)
	//fmt.Printf("ispis start")
	if err != nil {
		// If an error occurs other than record not found, return the error.
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ShoppingCart{}, err
		}

		// Initialize a new cart if it doesn't exist.
		newCart := &model.ShoppingCart{
			TouristId:  touristId,
			OrderItems: []model.OrderItem{},
			Total:      0,
		}

		// Create the new cart.
		createdCart, err := service.ShoppingCartRepo.Create(newCart)
		if err != nil {
			//fmt.Printf("ispis 1")
			return model.ShoppingCart{}, err
		}

		// Return the newly created cart.
		//fmt.Printf("ispis 2")
		return createdCart, nil
	}

	// Return the existing cart.
	//fmt.Printf("ispis 3")
	return cart, nil
}

func (service *ShoppingCartService) RemoveOrderItem(cartId int, tourId int) (*model.ShoppingCart, error) {
	// Retrieve the shopping cart by ID.
	cart, err := service.ShoppingCartRepo.Get(cartId)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve shopping cart: %w", err)
	}

	// Remove the order item by tour ID.
	cart.RemoveOrderItem(tourId)

	// Update the shopping cart in the repository.
	updatedCart, err := service.ShoppingCartRepo.Update(&cart)
	if err != nil {
		return nil, fmt.Errorf("failed to update shopping cart: %w", err)
	}

	// Map the updated cart to a DTO and return.
	return updatedCart, nil
}

func (service *ShoppingCartService) Update(cart *model.ShoppingCart) (*model.ShoppingCart, error) {

	// Update the shopping cart in the repository.
	updatedCart, err := service.ShoppingCartRepo.Update(cart)
	if err != nil {
		return nil, fmt.Errorf("failed to update shopping cart: %w", err)
	}

	// Map the updated cart to a DTO and return.
	return updatedCart, nil
}

func (service *ShoppingCartService) Purchase(cartId int) (model.ShoppingCart, error) {
	// Retrieve the shopping cart by ID
	cart, err := service.ShoppingCartRepo.Get(cartId)
	if err != nil {
		return model.ShoppingCart{}, err // Handle error if the cart is not found
	}

	for _, item := range cart.OrderItems {
		// For simplification, directly creating TourPurchaseToken
		token := model.NewTourPurchaseToken(cart.TouristId, item.IdTour)

		// Assuming Create function returns (*model.TourPurchaseToken, error)
		_, err := service.TokenRepo.Create(token)
		if err != nil {
			return model.ShoppingCart{}, err // Handle error during token creation
		}

		// Remove the order item from the cart (implementation depends on your model's method)
		// Assuming cart.RemoveOrderItem updates the cart's OrderItems slice
		cart.RemoveOrderItem(item.IdTour)
	}

	// Update the cart after processing its items
	_, err = service.ShoppingCartRepo.Update(&cart)
	if err != nil {
		return model.ShoppingCart{}, err
	}

	// Return the updated cart
	return cart, nil
}
