package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {
	//customer
	beego.GlobalControllerRouter["bee-example-api/controllers:CustomerController"] = append(beego.GlobalControllerRouter["bee-example-api/controllers:CustomerController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee-example-api/controllers:CustomerController"] = append(beego.GlobalControllerRouter["bee-example-api/controllers:CustomerController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee-example-api/controllers:CustomerController"] = append(beego.GlobalControllerRouter["bee-example-api/controllers:CustomerController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee-example-api/controllers:CustomerController"] = append(beego.GlobalControllerRouter["bee-example-api/controllers:CustomerController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee-example-api/controllers:CustomerController"] = append(beego.GlobalControllerRouter["bee-example-api/controllers:CustomerController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	//category
	beego.GlobalControllerRouter["bee-example-api/controllers:CategoryController"] = append(beego.GlobalControllerRouter["bee-example-api/controllers:CategoryController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee-example-api/controllers:CategoryController"] = append(beego.GlobalControllerRouter["bee-example-api/controllers:CategoryController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee-example-api/controllers:CategoryController"] = append(beego.GlobalControllerRouter["bee-example-api/controllers:CategoryController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee-example-api/controllers:CategoryController"] = append(beego.GlobalControllerRouter["bee-example-api/controllers:CategoryController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee-example-api/controllers:CategoryController"] = append(beego.GlobalControllerRouter["bee-example-api/controllers:CategoryController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	//product
	beego.GlobalControllerRouter["bee-example-api/controllers:ProductController"] = append(beego.GlobalControllerRouter["bee-example-api/controllers:ProductController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee-example-api/controllers:ProductController"] = append(beego.GlobalControllerRouter["bee-example-api/controllers:ProductController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee-example-api/controllers:ProductController"] = append(beego.GlobalControllerRouter["bee-example-api/controllers:ProductController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee-example-api/controllers:ProductController"] = append(beego.GlobalControllerRouter["bee-example-api/controllers:ProductController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee-example-api/controllers:ProductController"] = append(beego.GlobalControllerRouter["bee-example-api/controllers:ProductController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
}
