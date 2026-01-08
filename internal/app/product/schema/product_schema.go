package schema

import "time"

// product shcmea represetation tabel in database
type Product struct {
	ID            int       `json:"id"`
	ProductName   string    `json:"product_name"`
	Price         float64   `json:"price"`
	Stock         int       `json:"stock"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	CreatedAtUnix int64     `json:"created_at_unix"`
	UpdatedAt     time.Time `json:"updated_at"`
	UpdatedAtUnix int64     `json:"updated_at_unix"`
}

type ProductRequest struct {
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
}

type ProductResponse struct {
	ID          int       `json:"id"`
	ProductName string    `json:"product_name"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
