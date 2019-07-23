package routes

import (
	"github.com/betterde/ects/internal/middleware"
	"github.com/betterde/ects/web"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
)

func Register(app *iris.Application) {
	// SPA 路由刷新404 处理
	app.OnErrorCode(404, func(ctx iris.Context) {
		ctx.StatusCode(200)
		if err := ctx.View("index.html"); err != nil {
			log.Println(err)
		}
	})

	mvc.Configure(app.Party("/websocket"), registerWebSocket)

	// 接口路由
	mvc.Configure(app.PartyFunc("/api", func(api iris.Party) {
		// 鉴权路由
		mvc.Configure(api.Party("/auth"), authentication)

		api.Use(middleware.JWTHandler.Serve)

		// 节点管理路由
		mvc.Configure(api.Party("/node"), registerNode)

		// 任务流水线
		mvc.Configure(api.Party("/pipeline"), registerPipeline)

		// 任务
		mvc.Configure(api.Party("/task"), registerTask)

		// 组织路由
		mvc.Configure(api.PartyFunc("/organization", func(org iris.Party) {
			// 用户路由
			mvc.Configure(org.Party("/user"), registerUser)
			// 团队路由
			mvc.Configure(org.Party("/team"), registerTeam)
		}))

		// 角色路由
		mvc.Configure(api.Party("/role"), registerRole)

		// 权限路由
		mvc.Configure(api.Party("/permission"), registerPermission)

		// 日志路由
		mvc.Configure(api.Party("/log"), registerLog)

		// 账户路由
		mvc.Configure(api.PartyFunc("/account", func(account iris.Party) {
			// 用户个人信息
			mvc.Configure(account.Party("/profile"), registerProfile)
		}))
	}))

	// 压缩前端资源
	app.Use(iris.Gzip)

	app.RegisterView(iris.HTML("./web/dist", ".html").Binary(web.Asset, web.AssetNames))

	// 页面路由
	app.Get("/", func(ctx iris.Context) {
		if err := ctx.View("index.html"); err != nil {
			log.Println(err)
		}
	})

	assetHandler := iris.StaticEmbeddedHandler("./web/dist", web.Asset, web.AssetNames, false)
	app.SPA(assetHandler).AddIndexName("index.html")
}
