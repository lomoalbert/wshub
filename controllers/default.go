package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "home.html"
}

type VideoController struct {
	beego.Controller
}

func (c *VideoController) Get() {
	c.TplName = "video.html"
}