// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"time"
)

type Job struct {
	ID            int64     `json:"id"`
	Description   string    `json:"description"`
	ClientName    string    `json:"client_name"`
	ClientAddress string    `json:"client_address"`
	ClientEmail   string    `json:"client_email"`
	Price         int64     `json:"price"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	CreatedAt     time.Time `json:"created_at"`
}
