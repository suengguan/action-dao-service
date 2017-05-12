// @APIVersion 1.0.0
// @Title action_dao_service API
// @Description action_dao_service only serve the ACTION_T
// @Contact qsg@corex-tek.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"action-dao-service/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/action",
			beego.NSInclude(
				&controllers.ActionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
