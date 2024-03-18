package repo

import (
	"errors"

	"gorm.io/gorm"
	"payments.xws.com/model"
)

type CouponRepo struct {
	DatabaseConnection *gorm.DB
}

func (repo *CouponRepo) GetByCode(code string) (*model.Coupon, error) {
	coupon := &model.Coupon{}
	result := repo.DatabaseConnection.Where("code = ?", code).First(coupon)
	if result.Error != nil {
		return nil, result.Error
	}
	if coupon.IsExpired() {
		return nil, errors.New("coupon has expired")
	}
	return coupon, nil
}

func (repo *CouponRepo) Create(coupon *model.Coupon) (*model.Coupon, error) {
	coupon.GenerateCode()
	err := repo.DatabaseConnection.Create(coupon).Error
	if err != nil {
		return nil, err
	}
	return coupon, nil
}

func (repo *CouponRepo) GetByAuthorId(authorId int) ([]*model.Coupon, error) {
	var coupons []*model.Coupon
	result := repo.DatabaseConnection.Where("author_id = ?", authorId).Find(&coupons)
	if result.Error != nil {
		return nil, result.Error
	}
	return coupons, nil
}

func (repo *CouponRepo) Delete(id int) error {
	result := repo.DatabaseConnection.Delete(&model.Coupon{}, id)
	return result.Error
}

func (repo *CouponRepo) Update(coupon *model.Coupon) (*model.Coupon, error) {
	result := repo.DatabaseConnection.Save(coupon)
	if result.Error != nil {
		return nil, result.Error
	}
	return coupon, nil
}
