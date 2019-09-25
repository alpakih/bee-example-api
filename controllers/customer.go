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

type CustomerController struct {
	beego.Controller
}

func (c *CustomerController) Post() {
	var cDto dto.CreateCustomerDto
	errMap := json.Unmarshal(c.Ctx.Input.RequestBody, &cDto)
	if errMap != nil {
		c.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Data["json"] = response.BadRequest()
		c.ServeJSON()
		return
	}

	res, errMessage := dto.ValidateCustomerDto(cDto)
	if !res  {
		c.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Data["json"] = response.ValidationError(errMessage)
		c.ServeJSON()
		return
	}

	result, errInsert := services.StoreCustomer(cDto)
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

func (c *CustomerController) Get() {
	uid, _ := strconv.Atoi(c.GetString(":uid"))

	result, errGet := services.GetCustomerById(uid)

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

func (c *CustomerController) GetAll() {
	result, errGetAll := services.GetAllCustomer()
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

func (c *CustomerController) Put() {
	uid, _ := strconv.Atoi(c.GetString(":uid"))
	var cDto dto.UpdateCustomerDto

	errMap := json.Unmarshal(c.Ctx.Input.RequestBody, &cDto)

	if errMap != nil {
		c.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Data["json"] = response.BadRequest()
		c.ServeJSON()
		return
	}

	result, errUpdate := services.UpdateCustomer(uid, cDto)

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

func (c *CustomerController) Delete() {
	uid, _ := strconv.Atoi(c.GetString(":uid"))
	result, err := services.DeleteCustomer(uid)

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
