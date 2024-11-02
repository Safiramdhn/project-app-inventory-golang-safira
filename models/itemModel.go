package models

type Item struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Category   Category   `json:"category"`
	Location   Location   `json:"location"`
	Quantity   int        `json:"quantity"`
	Price      int        `json:"price"`
	Pagination Pagination `json:"pagination,omitempty"`
}
