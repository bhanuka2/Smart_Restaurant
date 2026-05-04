package services

import (
    "smart-restaurant/models"
    "smart-restaurant/repository"
)

func GenerateOffers() []models.Offer {
    products := repository.GetAllProducts()
    var offers []models.Offer

    for _, p := range products {
        if p.Stock < 5 {
            offer := models.Offer{
                ProductID: p.ID,
                Discount:  20,
                Reason:    "Low stock clearance",
            }
            offers = append(offers, offer)

            p.Discount = 20
            repository.UpdateProduct(p)
        }
    }

    return offers
}