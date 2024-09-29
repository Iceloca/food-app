package models

import productv1 "github.com/r1nb0/protos/gen/go/product"

type Product struct {
	ID            string
	Name          string
	Description   string
	Category      CategoryProduct
	Price         float32
	ImageURL      string
	IsBestProduct bool
}

type CreateProductDTO struct {
	Name          string
	Description   string
	Category      CategoryProduct
	Price         float32
	ImageURL      string
	IsBestProduct bool
}

type UpdateProductDTO struct {
	ID            string
	Name          string
	Description   string
	Category      CategoryProduct
	Price         float32
	ImageURL      string
	IsBestProduct bool
}

func (dto Product) MapToGRPCProduct() *productv1.Product {
	return &productv1.Product{
		Id:            dto.ID,
		Name:          dto.Name,
		Description:   dto.Description,
		Price:         dto.Price,
		ImageURL:      dto.ImageURL,
		IsBestProduct: dto.IsBestProduct,
		Category: &productv1.Category{
			Name:     dto.Category.Name,
			ImageURL: dto.Category.ImageURL,
		},
	}
}

func NewUpdateProductDTOFromGRPC(req *productv1.UpdateRequest) UpdateProductDTO {
	return UpdateProductDTO{
		ID:            req.Id,
		Name:          req.Name,
		Description:   req.Description,
		Price:         req.Price,
		ImageURL:      req.ImageURL,
		IsBestProduct: req.IsBestProduct,
		Category: CategoryProduct{
			Name:     req.Category.Name,
			ImageURL: req.Category.ImageURL,
		},
	}
}

func NewCreateProductDTOFromGRPC(req *productv1.CreateRequest) CreateProductDTO {
	return CreateProductDTO{
		Name:          req.Name,
		Description:   req.Description,
		Price:         req.Price,
		ImageURL:      req.ImageURL,
		IsBestProduct: req.IsBestProduct,
		Category: CategoryProduct{
			Name:     req.Category.Name,
			ImageURL: req.Category.ImageURL,
		},
	}
}
