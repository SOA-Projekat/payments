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

type ShoppingCartHandler struct {
	ShoppingCartService *service.ShoppingCartService
}

func (handler *ShoppingCartHandler) GetByUserId(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	idnum, _ := strconv.Atoi(id)
	log.Printf("Cart sa id-em %s", id)
	cart, err := handler.ShoppingCartService.GetByUserId(idnum)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(cart)
}

func (handler *ShoppingCartHandler) RemoveOrderItem(writer http.ResponseWriter, req *http.Request) {
	cartId, _ := strconv.Atoi(mux.Vars(req)["cartId"])
	tourId, _ := strconv.Atoi(mux.Vars(req)["tourId"])

	cart, err := handler.ShoppingCartService.RemoveOrderItem(cartId, tourId)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(cart)
}

func (handler *ShoppingCartHandler) Update(writer http.ResponseWriter, req *http.Request) {

	var obj model.ShoppingCart
	if err := json.NewDecoder(req.Body).Decode(&obj); err != nil {
		http.Error(writer, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	cart, err := handler.ShoppingCartService.Update(&obj)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(cart)
}

func (handler *ShoppingCartHandler) Purchase(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["cartId"]
	idnum, _ := strconv.Atoi(id)
	log.Printf("Cart sa id-em %s", id)
	cart, err := handler.ShoppingCartService.Purchase(idnum)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(cart)
}
