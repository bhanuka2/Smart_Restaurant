package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "smart-restaurant/services"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
    productID, _ := strconv.Atoi(r.URL.Query().Get("product_id"))
    qty, _ := strconv.Atoi(r.URL.Query().Get("qty"))

    order := services.CreateOrder(productID, qty)
    json.NewEncoder(w).Encode(order)
}