package main

import (
	"time"

	"official/models"
	_ "official/routers"

	"github.com/astaxie/beego"
	"official/utils"
	cache "github.com/patrickmn/go-cache"
	"github.com/astaxie/beego/logs"
)

func main() {
	models.Init()
	utils.Che = cache.New(60*time.Minute, 120*time.Minute)
	beego.SetLogger(logs.AdapterMultiFile,`{"filename":"./logs/official.log","separate":["emergency","critical", "error"]}`)
	beego.Run()
}
