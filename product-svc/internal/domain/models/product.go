package models

type Product struct {
	ID          string
	Name        string
	Description string
	Category    CategoryProduct
	Price       float64
	ImageURL    string
}

type CreateProductDTO struct {
	Name        string
	Description string
	Category    CategoryProduct
	Price       float64
	ImageURL    string
}

type UpdateProductDTO struct {
	ID          string
	Name        string
	Description string
	Category    CategoryProduct
	Price       float64
	ImageURL    string
}
