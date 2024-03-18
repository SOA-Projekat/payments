package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"payments.xws.com/model"
	"payments.xws.com/repo"
)

type CouponService struct {
	CouponRepo *repo.CouponRepo
}

func (service *CouponService) GetByCode(code string) (*model.Coupon, error) {
	return service.CouponRepo.GetByCode(code)
}

func (service *CouponService) Create(coupon *model.Coupon) (*model.Coupon, error) {
	return service.CouponRepo.Create(coupon)
}

func (service *CouponService) GetByAuthorId(authorId int) ([]*model.Coupon, error) {
	return service.CouponRepo.GetByAuthorId(authorId)
}

func (service *CouponService) Delete(id int) error {
	return service.CouponRepo.Delete(id)
}

func (service *CouponService) Update(coupon *model.Coupon) (*model.Coupon, error) {
	return service.CouponRepo.Update(coupon)
}

func (service *CouponService) CheckCoupon(code string, tourId int) (*model.Coupon, error) {

	if code == "" {
		return nil, errors.New("code is empty")
	}

	coupon, err := service.CouponRepo.GetByCode(code)

	if err != nil {
		return nil, errors.New("coupon is invalid or expired")
	}

	if coupon.TourId == -1 {

		url := fmt.Sprintf("https://localhost:44333/api/administration/tour/checkTour/%d/%d", coupon.AuthorId, tourId)

		// Make the GET request
		resp, err := http.Get(url)
		if err != nil {
			// Handle error (e.g., network error)
			fmt.Println("Error making GET request:", err)
			return nil, errors.New("error making GET request")
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			// Handle error (e.g., error reading the response)
			fmt.Println("Error reading response body:", err)
			return nil, errors.New("error reading response body")
		}

		// Assuming the response is a boolean wrapped in a JSON object
		// Define a struct to match the expected JSON structure
		var value bool

		// Unmarshal the JSON response into the struct
		if err := json.Unmarshal(body, &value); err != nil {
			// Handle JSON parsing error
			fmt.Println("Error parsing JSON response:", err)
			return nil, errors.New("error parsing JSON response")
		}

		if !value {
			return nil, errors.New("entered coupon does not belong to author of this tour")
		}
	} else if coupon.TourId != tourId {
		return nil, errors.New("entered coupon is not for this tour")
	}

	return coupon, nil

}
