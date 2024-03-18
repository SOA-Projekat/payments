package service

import (
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
