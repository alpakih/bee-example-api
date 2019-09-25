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

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Post() {
	var cDto dto.CategoryDto
	errMap := json.Unmarshal(c.Ctx.Input.RequestBody, &cDto)
	if errMap != nil {
		c.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Data["json"] = response.BadRequest()
		c.ServeJSON()
		return
	}

	res, errMessage := dto.ValidateCategoryDto(cDto)
	if !res  {
		c.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Data["json"] = response.ValidationError(errMessage)
		c.ServeJSON()
		return
	}

	result, errInsert := services.StoreCategory(cDto)
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

func (c *CategoryController) Get() {
	uid, _ := strconv.Atoi(c.GetString(":uid"))

	result, errGet := services.GetCategoryById(uid)

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

func (c *CategoryController) GetAll() {
	result, errGetAll := services.GetAllCategory()
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

func (c *CategoryController) Put() {
	uid, _ := strconv.Atoi(c.GetString(":uid"))
	var cDto dto.CategoryDto

	errMap := json.Unmarshal(c.Ctx.Input.RequestBody, &cDto)

	if errMap != nil {
		c.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Data["json"] = response.BadRequest()
		c.ServeJSON()
		return
	}

	result, errUpdate := services.UpdateCategory(uid, cDto)

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

func (c *CategoryController) Delete() {
	uid, _ := strconv.Atoi(c.GetString(":uid"))
	result, err := services.DeleteCategory(uid)

	if err != nil && !result {
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
