package schema

// product shcmea represetation tabel in database
type Product struct {
	ID            string  `json:"id"`
	ProductName   string  `json:"product_name"`
	Price         float64 `json:"price"`
	Stock         int     `json:"stock"`
	Description   string  `json:"description"`
	CreatedAt     string  `json:"created_at"`
	CreatedAtUnix int64   `json:"created_at_unix"`
	UpdatedAt     string  `json:"updated_at"`
	UpdatedAtUnix int64   `json:"updated_at_unix"`
}
