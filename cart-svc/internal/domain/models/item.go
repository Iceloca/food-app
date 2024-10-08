package models

import cartv1 "github.com/r1nb0/protos/gen/go/cart"

type Item struct {
	Product Product
	Count   int64
}

func NewItemFromGRPC(item *cartv1.Item) Item {
	return Item{
		Product: NewProductCreateFromGRPC(item.Product),
		Count:   item.Count,
	}
}

func (i *Item) MapItemToGRPC() *cartv1.Item {
	return &cartv1.Item{
		Product: i.Product.MapProductToGRPC(),
		Count:   i.Count,
	}
}
