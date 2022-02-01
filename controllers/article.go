package controllers

import (
	"fmt"
	"gin/Dtos"
	"gin/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

var repository *services.ArticleRepository

// InitArticleController 初始化博文Controller
func InitArticleController() {
	repository = services.NewArticleRepository(true)
}

func AddArticle(c *gin.Context) {
	var articleDto Dtos.ArticleAddDto

	if err := c.ShouldBind(&articleDto); err != nil {
		c.JSON(http.StatusBadRequest, Dtos.ErrorDto{
			Message:          "Bind Model Error",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	entity := articleDto.ToEntity(1)
	repository.AddArticle(entity)

	c.JSON(http.StatusCreated, Dtos.ParseArticleEntity(entity))
}

// GetArticle 添加博文
func GetArticle(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if !repository.ArticleExists(uint(id)) {
		c.JSON(http.StatusNotFound, Dtos.ErrorDto{
			Message:          fmt.Sprintf("Article %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	article := repository.GetArticle(uint(id))
	// 转换为Dto
	c.JSON(http.StatusOK, Dtos.ParseArticleEntity(article))
}

func GetArticles(c *gin.Context) {
	articles := repository.GetArticles()

	// 转换为Dto
	articleDtos := make([]Dtos.ArticleGetDto, len(articles))
	for _, article := range articles {
		articleDtos = append(articleDtos, *Dtos.ParseArticleEntity(&article))
	}

	c.JSON(http.StatusOK, articles)
}

func UpdateArticle(c *gin.Context) {

	c.JSON(http.StatusOK, "articleDtos")
}

func DeleteArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "AddBlog"})
}
