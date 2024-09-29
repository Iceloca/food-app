package models

import productv1 "github.com/r1nb0/protos/gen/go/product"

type CategoryProduct struct {
	Name     string
	ImageURL string
}

func (dto *CategoryProduct) MapToGRPCCategory() *productv1.Category {
	return &productv1.Category{
		Name:     dto.Name,
		ImageURL: dto.ImageURL,
	}
}
