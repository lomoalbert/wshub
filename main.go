package main

import (
	_ "github.com/lomoalbert/wshub/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.EnableAdmin = true
	beego.Run()
}

