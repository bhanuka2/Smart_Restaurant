package controllers

import (
    "encoding/json"
    "net/http"
    "smart-restaurant/repository"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
    products := repository.GetAllProducts()
    json.NewEncoder(w).Encode(products)
}