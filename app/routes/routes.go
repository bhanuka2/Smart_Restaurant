package routes

import (
    "net/http"
    "smart-restaurant/controllers"
)

func RegisterRoutes() {
    http.HandleFunc("/products", controllers.GetProducts)
    http.HandleFunc("/order", controllers.CreateOrder)
}