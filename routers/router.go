package routers

import (
	"github.com/astaxie/beego"
	"official/controllers"
	"official/controllers/api"
)

func init() {
	// 默认登录
	beego.Router("/login", &controllers.LoginController{}, "*:LoginIn")
	beego.Router("/login_out", &controllers.LoginController{}, "*:LoginOut")
	beego.Router("/no_auth", &controllers.LoginController{}, "*:NoAuth")

	beego.Router("/", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")

	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.AdminController{})
	beego.AutoRouter(&controllers.UserController{})

	beego.AutoRouter(&controllers.NewsController{})//新闻资讯
	beego.AutoRouter(&controllers.UploadController{})//文件上传
	beego.AutoRouter(&controllers.CasesController{})//案例
	beego.AutoRouter(&controllers.LeaderController{})//项目负责人
	beego.AutoRouter(&controllers.BannerController{})//Banner
	beego.AutoRouter(&controllers.ContractController{})//Contract

	beego.Router("/api/news/?:cate", &api.NewsApiController{}, "get:NewsList")
	beego.Router("/api/news_detail/?:id", &api.NewsApiController{}, "get:NewsDetail")
	beego.Router("/api/trend", &api.TrendApiController{}, "get:Trend")
	beego.Router("/api/trend_cover", &api.TrendApiController{}, "get:TrendCover")
	beego.Router("/api/case", &api.CasesApiController{}, "get:CasesList")
	beego.Router("/api/case_detail/?:id", &api.CasesApiController{}, "get:CasesDetailByLeaderId")
	beego.Router("/api/case_info/?:caseKey", &api.CasesApiController{}, "get:CasesDetailByType")
	beego.Router("/api/leader", &api.CasesApiController{}, "get:LeaderList")
	beego.Router("/api/contract", &api.ContractApiController{}, "get:ContractList")
	beego.Router("/api/banner", &api.BannerApiController{}, "get:BannerList")
}
