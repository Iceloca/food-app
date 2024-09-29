package models

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
