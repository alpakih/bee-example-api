package dto

type ProductDto struct {
	Name        string `json:"name" valid:"Required"`
	Description string `json:"description" valid:"Required"`
	Price       int64  `json:"price" valid:"Required"`
	CategoryId  int    `json:"category_id" valid:"Required"`
}
