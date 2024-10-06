package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	InStock     = 1 << 0 // 00001
	OnSale      = 1 << 1 // 00010
	Discounted  = 1 << 2 // 00100
	ExpressShip = 1 << 3 // 01000
	Refurbished = 1 << 4 // 10000
)

type ProductFeature struct {
	ID             uint  `gorm:"primaryKey"`
	UserID         int   `json:"user_id"`
	ProductID      int   `json:"product_id"`
	PackedFeatures uint8 `json:"packed_features"`
}

type Features struct {
	InStock     bool `json:"in_stock"`
	OnSale      bool `json:"on_sale"`
	Discounted  bool `json:"discounted"`
	ExpressShip bool `json:"express_ship"`
	Refurbished bool `json:"refurbished"`
}

type InsertBody struct {
	UserID    int      `json:"user_id"`
	ProductID int      `json:"product_id"`
	Features  Features `json:"features"`
}

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database:", err)
	}
	db.AutoMigrate(&ProductFeature{})
}

func packProduct(features Features) uint8 {
	var packed uint8
	if features.InStock {
		packed |= InStock
	}
	if features.OnSale {
		packed |= OnSale
	}
	if features.Discounted {
		packed |= Discounted
	}
	if features.ExpressShip {
		packed |= ExpressShip
	}
	if features.Refurbished {
		packed |= Refurbished
	}
	return packed
}

func createProductFeature(w http.ResponseWriter, r *http.Request) {
	var input InsertBody
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	packedFeatures := packProduct(input.Features)

	productFeature := ProductFeature{
		UserID:         input.UserID,
		ProductID:      input.ProductID,
		PackedFeatures: packedFeatures,
	}
	db.Create(&productFeature)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(productFeature)
}

func updateProductFeature(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var input Features
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var productFeature ProductFeature
	if err := db.First(&productFeature, id).Error; err != nil {
		http.Error(w, "ProductFeature not found", http.StatusNotFound)
		return
	}

	productFeature.PackedFeatures = packProduct(input)
	db.Save(&productFeature)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productFeature)
}

func getProductFeature(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var productFeature ProductFeature
	if err := db.First(&productFeature, id).Error; err != nil {
		http.Error(w, "ProductFeature not found", http.StatusNotFound)
		return
	}

	features := Features{
		InStock:     productFeature.PackedFeatures&InStock > 0,
		OnSale:      productFeature.PackedFeatures&OnSale > 0,
		Discounted:  productFeature.PackedFeatures&Discounted > 0,
		ExpressShip: productFeature.PackedFeatures&ExpressShip > 0,
		Refurbished: productFeature.PackedFeatures&Refurbished > 0,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(features)
}

func main() {
	initDB()

	http.HandleFunc("/product_features", createProductFeature)        // POST
	http.HandleFunc("/product_features/update", updateProductFeature) // PUT
	http.HandleFunc("/product_features", getProductFeature)           // GET

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
