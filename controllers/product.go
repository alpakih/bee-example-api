package controllers

import (
	"bee-example-api/dto"
	"bee-example-api/response"
	"bee-example-api/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"net/http"
	"strconv"
)

type ProductController struct {
	beego.Controller
}

func (c *ProductController) Post() {
	var cDto dto.ProductDto
	errMap := json.Unmarshal(c.Ctx.Input.RequestBody, &cDto)

	if errMap != nil {
		c.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Data["json"] = response.BadRequest()
		c.ServeJSON()
		return
	}
	result, errInsert := services.StoreProduct(cDto)
	if errInsert != nil {
		c.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Data["json"] = response.BadRequest()
		c.ServeJSON()
		return
	}
	c.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Data["json"] = response.Single(result)
	c.ServeJSON()
	return
}
func (c *ProductController) Get() {
	uid, _ := strconv.Atoi(c.GetString(":uid"))

	result, errGet := services.GetProductById(uid)

	if errGet != nil {
		c.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
		c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
		c.Data["json"] = response.NotFound()
		c.ServeJSON()
		return
	}
	c.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Data["json"] = response.Single(result)
	c.ServeJSON()

}

func (c *ProductController) GetAll() {
	result, errGetAll := services.GetAllProduct()
	if errGetAll != nil {
		c.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
		c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
		c.Data["json"] = response.NotFound()
		c.ServeJSON()
		return
	}

	c.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Data["json"] = response.List(result)
	c.ServeJSON()
}

func (c *ProductController) Put() {
	uid, _ := strconv.Atoi(c.GetString(":uid"))
	var cDto dto.ProductDto

	errMap := json.Unmarshal(c.Ctx.Input.RequestBody, &cDto)

	if errMap != nil {
		c.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Data["json"] = response.BadRequest()
		c.ServeJSON()
		return
	}

	result, errUpdate := services.UpdateProduct(uid, cDto)

	if errUpdate != nil {
		c.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Data["json"] = response.BadRequest()
		c.ServeJSON()
		return
	}
	c.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Data["json"] = response.Single(result)
	c.ServeJSON()
	return
}

func (c *ProductController) Delete() {
	uid, _ := strconv.Atoi(c.GetString(":uid"))
	result, err := services.DeleteProduct(uid)

	if err != nil || !result {
		c.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Data["json"] = response.BadRequest()
		c.ServeJSON()
		return
	}
	c.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Data["json"] = response.Success("Success deleted")
	c.ServeJSON()
	return
}
