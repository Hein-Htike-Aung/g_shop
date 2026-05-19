package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	_ "yixiang.co/go-mall/docs"
	"yixiang.co/go-mall/middleware"
	"yixiang.co/go-mall/pkg/upload"
	"yixiang.co/go-mall/routers/admin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// register admin routes
	admin.RegisterSystemRouters(r)
	admin.RegisterShopRouters(r)
	admin.RegisterToolsRouters(r)
	admin.RegisterWeixinRouters(r)

	// register API routes
	//api.RegisterApiRouters(r)

	return r
}
