package routers

import (
	"book-mgr-backend/handler/admin"
	"book-mgr-backend/handler/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type App struct {
}

func (a *App) RunServer() {
	r := gin.Default()

	r.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(http.StatusOK)
			return
		}
		context.Next()
	})

	r.GET("books")

	adminGroup := r.Group("/api/admin/v1")
	{
		adminGroup.POST("login", admin.HandleAdminLogin)

		adminGroup.GET("book", admin.HandleGetAllBooks_Admin)
		adminGroup.POST("book", admin.HandleAddBook_Admin)
		adminGroup.PUT("book", admin.HandleUpdateBook_Admin)
		adminGroup.DELETE("book", admin.HandleDeleteBook_Admin)

		adminGroup.GET("user", admin.HandleGetAllUsers_Admin)
	}

	userGroup := r.Group("/api/user/v1")
	{
		userGroup.POST("login")
		userGroup.GET("book", user.HandleGetAllBooks_User)
		userGroup.GET("history", user.HandleGetAllMyBorrowed_User)
	}

	if err := r.Run("localhost:7001"); err != nil {
		log.Panicln("端口可能已被占用 服务器启动失败" + err.Error())
	}
}
