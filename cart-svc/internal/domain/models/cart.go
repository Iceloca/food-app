package models

import cartv1 "github.com/r1nb0/protos/gen/go/cart"

type Cart struct {
	ID         string
	UserID     int64
	Items      []Item
	TotalPrice float32
}

type CartCreate struct {
	UserID     int64
	Items      []Item
	TotalPrice float32
}

func NewCartCreateFromGRPC(cart *cartv1.CreateRequest) CartCreate {
	var items []Item
	for _, protoItem := range cart.Items {
		item := NewItemFromGRPC(protoItem)
		items = append(items, item)
	}
	return CartCreate{
		UserID:     cart.UserId,
		Items:      items,
		TotalPrice: cart.TotalPrice,
	}
}

func (c *Cart) MapCartToGRPC() *cartv1.Cart {
	var items []*cartv1.Item

	for _, item := range c.Items {
		protoItem := item.MapItemToGRPC()
		items = append(items, protoItem)
	}

	return &cartv1.Cart{
		Id:         c.ID,
		UserId:     c.UserID,
		Items:      items,
		TotalPrice: c.TotalPrice,
	}
}

func NewCartFromGRPC(cart *cartv1.Cart) Cart {
	var items []Item
	for _, protoItem := range cart.Items {
		item := NewItemFromGRPC(protoItem)
		items = append(items, item)
	}
	return Cart{
		ID:         cart.Id,
		UserID:     cart.UserId,
		Items:      items,
		TotalPrice: cart.TotalPrice,
	}
}
