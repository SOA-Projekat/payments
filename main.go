package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"payments.xws.com/handler"
	"payments.xws.com/model"
	"payments.xws.com/repo"
	"payments.xws.com/service"
)

func initDB() *gorm.DB {

	dsn := "user=postgres password=super dbname=payments host=localhost port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.ShoppingCart{})
	database.AutoMigrate(&model.TourPurchaseToken{})
	database.AutoMigrate(&model.Coupon{})
	return database
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	//Repositories
	cartRepo := &repo.ShoppingCartRepo{DatabaseConnection: database}
	tokenRepo := &repo.TourPurchaseTokenRepository{DatabaseConnection: database}
	couponRepo := &repo.CouponRepo{DatabaseConnection: database}

	//Services
	cartService := &service.ShoppingCartService{
		ShoppingCartRepo: cartRepo,
		TokenRepo:        tokenRepo,
	}
	tokenService := &service.TourPurchaseTokenService{TokenRepository: tokenRepo}
	couponService := &service.CouponService{CouponRepo: couponRepo}
	//Handlers
	cartHandler := &handler.ShoppingCartHandler{ShoppingCartService: cartService}
	tokenHandler := &handler.TourPurchaseTokenHandler{TokenHandler: tokenService}
	couponHandler := &handler.CouponHandler{CouponService: couponService}
	// Router setup
	router := mux.NewRouter().StrictSlash(true)

	// Routes for carts
	router.HandleFunc("/shoppingcart/{id}", cartHandler.GetByUserId).Methods("GET")
	router.HandleFunc("/shoppingcart/purchase/{cartId}", cartHandler.Purchase).Methods("PUT")
	router.HandleFunc("/shoppingcart/{cartId}/{tourId}", cartHandler.RemoveOrderItem).Methods("PUT")
	router.HandleFunc("/shoppingcart/update", cartHandler.Update).Methods("PUT")

	// Routes for tokens
	router.HandleFunc("/tokens", tokenHandler.GetAll).Methods("GET")
	router.HandleFunc("/tokens/{touristId}", tokenHandler.GetAllByTourist).Methods("GET")

	//Routes for coupons
	router.HandleFunc("/authoring/coupon/getByCode", couponHandler.GetByCode).Methods("GET")
	router.HandleFunc("/authoring/coupon/{authorId}", couponHandler.GetByAuthorId).Methods("GET")
	router.HandleFunc("/authoring/coupon/{id}", couponHandler.Update).Methods("PUT")
	router.HandleFunc("/authoring/coupon/{id}", couponHandler.Delete).Methods("DELETE")
	router.HandleFunc("/authoring/coupon", couponHandler.Create).Methods("POST")
	router.HandleFunc("/checkCoupon", couponHandler.CheckCoupon).Methods("GET")

	// CORS setup
	permittedHeaders := handlers.AllowedHeaders([]string{"Requested-With", "Content-Type", "Authorization"})
	permittedOrigins := handlers.AllowedOrigins([]string{"*"})
	permittedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	// Start server
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8082", handlers.CORS(permittedHeaders, permittedOrigins, permittedMethods)(router)))
}
