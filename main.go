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

	err = database.AutoMigrate(&model.ShoppingCart{})
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

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

	//Services
	cartService := &service.ShoppingCartService{ShoppingCartRepo: cartRepo}

	//Handlers
	cartHandler := &handler.ShoppingCartHandler{ShoppingCartService: cartService}

	// Router setup
	router := mux.NewRouter().StrictSlash(true)

	// Routes for carts
	router.HandleFunc("/shoppingcart/{id}", cartHandler.GetByUserId).Methods("GET")
	router.HandleFunc("/shoppingcart/{cartId}/{tourId}", cartHandler.RemoveOrderItem).Methods("PUT")
	router.HandleFunc("/shoppingcart/update", cartHandler.RemoveOrderItem).Methods("PUT")

	// CORS setup
	permittedHeaders := handlers.AllowedHeaders([]string{"Requested-With", "Content-Type", "Authorization"})
	permittedOrigins := handlers.AllowedOrigins([]string{"*"})
	permittedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	// Start server
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8082", handlers.CORS(permittedHeaders, permittedOrigins, permittedMethods)(router)))
}
