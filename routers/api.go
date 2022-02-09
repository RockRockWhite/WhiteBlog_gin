package routers

import (
    "gin/controllers"
    "gin/middlewares"
    "github.com/gin-gonic/gin"
)

func InitApiRouter() *gin.Engine {
	// 初始化Controllers
	controllers.InitArticleController()
	controllers.InitUserController()

	router := gin.Default()

	// 配置中间件
	router.Use(middlewares.Logger())

	blog := router.Group("/article")
	{
		blog.GET("/:id", controllers.GetArticle)
		blog.GET("/", controllers.GetArticles)
		blog.POST("/", middlewares.JwtAuth(), controllers.AddArticle)
		blog.PUT("/", middlewares.JwtAuth(), controllers.UpdateArticle)
		blog.PATCH("/:id", middlewares.JwtAuth(), controllers.UpdateArticle)
		blog.DELETE("/:id", middlewares.JwtAuth(), controllers.DeleteArticle)
	}

	user := router.Group("/user")
	{
		user.GET("/:id", controllers.GetUser)
		user.POST("/", controllers.AddUser)
		user.PUT("/", controllers.PutUser)
		user.PATCH("/:id", controllers.PatchUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}

	token := router.Group("/token")
	{
		token.POST("", controllers.CreateToken)
	}

	return router
}
