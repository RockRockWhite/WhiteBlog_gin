package routers

import (
	"gin/controllers"
	"github.com/gin-gonic/gin"
)

func InitApiRouter() *gin.Engine {
	// 初始化Controllers
	controllers.InitArticleController()

	router := gin.Default()

	blog := router.Group("/article")
	{
		blog.GET("/:id", controllers.GetArticle)
		blog.GET("/", controllers.GetArticles)
		blog.POST("/", controllers.AddArticle)
		blog.PUT("/", controllers.UpdateArticle)
		blog.PATCH("/", controllers.UpdateArticle)
		blog.DELETE("/", controllers.DeleteArticle)
	}
	return router
}
