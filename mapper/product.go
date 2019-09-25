package mapper

type ProductMapper struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Price        int64  `json:"price"`
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}
