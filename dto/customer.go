package dto

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

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

func ValidateCustomerDto(u CreateCustomerDto) (res bool, dataErrorValidation interface{}) {
	valid := validation.Validation{}

	resultValidation, _ := valid.Valid(&u)
	if !resultValidation {
		merged := make(beego.M)
		for _, errRange := range valid.Errors {
			merged[errRange.Field] = errRange.Message
		}
		dataErrorValidation = merged
		res = false
		return res, dataErrorValidation
	}
	res = true
	return res, dataErrorValidation
}
