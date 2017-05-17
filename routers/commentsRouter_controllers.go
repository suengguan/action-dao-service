package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["dao-service/action-dao-service/controllers:ActionController"] = append(beego.GlobalControllerRouter["dao-service/action-dao-service/controllers:ActionController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["dao-service/action-dao-service/controllers:ActionController"] = append(beego.GlobalControllerRouter["dao-service/action-dao-service/controllers:ActionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/:userId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
