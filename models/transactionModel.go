package models

import "time"

type Transaction struct {
	ID              int        `json:"id"`
	Item            Item       `json:"item"`
	Quantity        int        `json:"quantity"`
	TotalPrice      int        `json:"total_price"`
	Timestamp       time.Time  `json:"timestamp"`
	TransactionType string     `json:"type"`
	AddedBy         int        `json:"added_by"`
	Description     string     `json:"description"`
	Pagination      Pagination `json:"pagination,omitempty"`
}
