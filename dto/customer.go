package dto

type CreateCustomerDto struct {
	Name    string `json:"name" valid:"Required"`
	Email   string `json:"email" valid:"Required"`
	Address string `json:"address" valid:"Required"`
	Phone   string `json:"phone" valid:"Required"`
}

type UpdateCustomerDto struct {
	Name    string `json:"name" valid:"Required"`
	Email   string `json:"email" valid:"Required"`
	Address string `json:"address" valid:"Required"`
	Phone   string `json:"phone" valid:"Required"`
}
