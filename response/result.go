package response

import (
	"net/http"
)

func Single(d Data) Response {
	meta := MetaComposer(http.StatusOK, "OK")
	return Response{Meta: meta, Data: d.Map()}
}

func List(d interface{}) Response {
	meta := MetaComposer(http.StatusOK, "OK")
	return Response{Meta: meta, Data: d}
}

func BadRequest() Response {
	meta := MetaComposer(http.StatusBadRequest, "Bad Request")
	return Response{Meta: meta, Data: nil}
}

func NotFound() Response {
	meta := MetaComposer(http.StatusNotFound, "Data Not Found")
	return Response{Meta: meta, Data: nil}
}

func Success(message string) Response {
	meta := MetaComposer(http.StatusOK, message)
	return Response{Meta: meta, Data: nil}
}

func ValidationError(d interface{}) Response {
	meta := MetaComposer(http.StatusBadRequest, "Validation Error")
	return Response{Meta: meta, Data: d}
}
