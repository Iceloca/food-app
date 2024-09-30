package models

import productv1 "github.com/r1nb0/protos/gen/go/product"

type Product struct {
	ID          int64
	Name        string
	Description string
	ImageURL    string
	Price       float32
	IsDailyRec  bool
	CategoryID  int64
}

type ProductCreate struct {
	Name        string
	Description string
	ImageURL    string
	Price       float32
	IsDailyRec  bool
	CategoryID  int64
}

func NewProductCreateFromGRPC(req *productv1.CreateRequest) ProductCreate {
	return ProductCreate{
		Name:        req.Name,
		Description: req.Description,
		ImageURL:    req.ImageUrl,
		Price:       req.Price,
		IsDailyRec:  req.IsDailyRec,
		CategoryID:  req.CategoryId,
	}
}

func NewProductFromGRPC(req *productv1.Product) Product {
	return Product{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		ImageURL:    req.ImageUrl,
		Price:       req.Price,
		IsDailyRec:  req.IsDailyRec,
		CategoryID:  req.CategoryId,
	}
}

func (p *Product) MapToGRPCProduct() *productv1.Product {
	return &productv1.Product{
		Id:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		ImageUrl:    p.ImageURL,
		Price:       p.Price,
		IsDailyRec:  p.IsDailyRec,
		CategoryId:  p.CategoryID,
	}
}
