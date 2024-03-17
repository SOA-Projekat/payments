package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"payments.xws.com/service"
)

type TourPurchaseTokenHandler struct {
	TokenHandler *service.TourPurchaseTokenService
}

func (handler *TourPurchaseTokenHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	tokens, err := handler.TokenHandler.GetAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tokens)
}

func (handler *TourPurchaseTokenHandler) GetAllByTourist(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["touristId"]
	idnum, _ := strconv.Atoi(id)
	cart, err := handler.TokenHandler.GetAllByTourist(idnum)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(cart)
}
