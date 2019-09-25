package dto

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type ProductDto struct {
	Name        string `json:"name" valid:"Required"`
	Description string `json:"description" valid:"Required"`
	Price       int64  `json:"price" valid:"Required"`
	CategoryId  int    `json:"category_id" valid:"Required"`
}

func ValidateProductDto(u ProductDto) (res bool, dataErrorValidation interface{}) {
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