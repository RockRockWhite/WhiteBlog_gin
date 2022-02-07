package routers

import (
	"gin/controllers"
	"gin/middlewares"
	"github.com/gin-gonic/gin"
)

func InitApiRouter() *gin.Engine {
	// 初始化Controllers
	controllers.InitArticleController()

	router := gin.Default()

	// 配置中间件
	router.Use(middlewares.Logger())

	blog := router.Group("/article")
	{
		blog.GET("/:id", controllers.GetArticle)
		blog.GET("/", controllers.GetArticles)
		blog.POST("/", controllers.AddArticle)
		blog.PUT("/", controllers.UpdateArticle)
		blog.PATCH("/:id", controllers.UpdateArticle)
		blog.DELETE("/:id", controllers.DeleteArticle)
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
		token.GET("/:id", controllers.GetToken)
	}

	return router
}
