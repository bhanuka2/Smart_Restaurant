package repository

import "smart-restaurant/models"

var products = []models.Product{
    {ID: 1, Name: "Burger", Price: 500, Stock: 3},
    {ID: 2, Name: "Pizza", Price: 1200, Stock: 15},
}

func GetAllProducts() []models.Product {
    return products
}

func UpdateProduct(updated models.Product) {
    for i, p := range products {
        if p.ID == updated.ID {
            products[i] = updated
        }
    }
}