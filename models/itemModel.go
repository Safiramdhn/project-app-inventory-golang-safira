package models

type Item struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CategoryId int    `json:"category_id"`
	LocationId int    `json:"location_id"`
	Quantity   int    `json:"quantity"`
	Price      int    `json:"price"`
}
