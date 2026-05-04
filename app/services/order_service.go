package services

import (
    "smart-restaurant/models"
    "smart-restaurant/repository"
)

func CreateOrder(productID int, qty int) models.Order {
    products := repository.GetAllProducts()

    for _, p := range products {
        if p.ID == productID {
            price := p.Price
            if p.Discount > 0 {
                price = price - (price * p.Discount / 100)
            }

            total := price * float64(qty)

            return models.Order{
                ProductID: productID,
                Quantity:  qty,
                Total:     total,
            }
        }
    }

    return models.Order{}
}