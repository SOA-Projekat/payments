package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"payments.xws.com/model"
	"payments.xws.com/service"
)

type CouponHandler struct {
	CouponService *service.CouponService
}

func (handler *CouponHandler) GetByCode(writer http.ResponseWriter, req *http.Request) {
	code := req.URL.Query().Get("code")
	log.Printf("Coupon with code: %s", code)

	coupon, err := handler.CouponService.GetByCode(code)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(coupon)
}

func (handler *CouponHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var coupon model.Coupon
	if err := json.NewDecoder(req.Body).Decode(&coupon); err != nil {
		http.Error(writer, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	createdCoupon, err := handler.CouponService.Create(&coupon)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(createdCoupon)
}

func (handler *CouponHandler) GetByAuthorId(writer http.ResponseWriter, req *http.Request) {
	authorIdStr := mux.Vars(req)["authorId"]
	authorId, _ := strconv.Atoi(authorIdStr)

	coupons, err := handler.CouponService.GetByAuthorId(authorId)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(coupons)
}

func (handler *CouponHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, _ := strconv.Atoi(idStr)

	err := handler.CouponService.Delete(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

func (handler *CouponHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var coupon model.Coupon
	if err := json.NewDecoder(req.Body).Decode(&coupon); err != nil {
		http.Error(writer, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	updatedCoupon, err := handler.CouponService.Update(&coupon)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(updatedCoupon)
}

func (handler *CouponHandler) CheckCoupon(writer http.ResponseWriter, req *http.Request) {
	code := req.URL.Query().Get("code")
	tourId, _ := strconv.Atoi(req.URL.Query().Get("tourId"))
	//log.Printf("Coupon with code: %s", code)

	coupon, err := handler.CouponService.CheckCoupon(code, tourId)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(coupon)
}
