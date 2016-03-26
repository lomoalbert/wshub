package routers

import (
	"github.com/lomoalbert/wshub/controllers"
	"github.com/astaxie/beego"
	"golang.org/x/net/websocket"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Handler("/ws", websocket.Handler(controllers.GroupChatServer))
	beego.Router("/video/", &controllers.VideoController{})
}
