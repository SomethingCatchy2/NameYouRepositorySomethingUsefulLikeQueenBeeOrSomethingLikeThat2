package routers

import (
	"goBee/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/about", &controllers.AboutController{})
	beego.Router("/aboat", &controllers.BoatController{})
	beego.Router("/contact", &controllers.ContactController{})
	beego.Router("/ukk", &controllers.UkkController{})
	beego.Router("/profile", &controllers.ProfileController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/admin/login", &controllers.LoginController{})
	beego.Router("/admin/logout", &controllers.LogoutController{})
	beego.Router("/admin/contacts/:id", &controllers.ContactDetailController{})
	beego.Router("/admin/contacts", &controllers.ContactListController{})

}
