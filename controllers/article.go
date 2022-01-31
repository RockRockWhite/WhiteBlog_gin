package controllers

import (
	"gin/Dtos"
	"gin/services"
	"github.com/gin-gonic/gin"
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
	if err := c.BindJSON(&articleDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	entity := articleDto.ToEntity(1)
	if _, err := repository.AddArticle(entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, Dtos.ParseArticleEntity(entity))
}

// GetArticle 添加博文
func GetArticle(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if !repository.ArticleExists(uint(id)) {
		c.Status(http.StatusNotFound)
		return
	}

	article, _ := repository.GetArticle(uint(id))
	// 转换为Dto
	c.JSON(http.StatusOK, Dtos.ParseArticleEntity(article))
}

func GetArticles(c *gin.Context) {
	articles, _ := repository.GetArticles()

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
