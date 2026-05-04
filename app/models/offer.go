package models

type Offer struct {
    ProductID int     `json:"product_id"`
    Discount  float64 `json:"discount"`
    Reason    string  `json:"reason"`
}