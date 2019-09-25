package dto

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type CategoryDto struct {
	Code string `json:"code" valid:"Required"`
	Name string `json:"name" valid:"Required"`
}

func (u *CategoryDto) Valid(v *validation.Validation) {
	//custom validation
}

func ValidateCategoryDto(u CategoryDto) (res bool, dataErrorValidation interface{}) {
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
