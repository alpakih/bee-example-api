package response

type Meta struct {
	Code		int 	`json:"code"`
	Message 	string 	`json:"message"`
}

func MetaComposer(code int, message string) Meta {
	meta := Meta{
		Code:    code,
		Message: message,
	}
	return meta
}