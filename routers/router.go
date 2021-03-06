package routers

import (
	"Kenai-plus/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &controllers.HelloController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/write", &controllers.WriterController{})
	beego.Router("/fileUpload", &controllers.FileUploadController{})
}
