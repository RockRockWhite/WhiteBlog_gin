package controllers

import (
	"fmt"
	"gin/dtos"
	"gin/services"
	"gin/utils"
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
	var articleDto dtos.ArticleAddDto

	if err := c.ShouldBind(&articleDto); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Bind Model Error",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	entity := articleDto.ToEntity(1)
	repository.AddArticle(entity)

	c.JSON(http.StatusCreated, dtos.ParseArticleEntity(entity))
}

// GetArticle 添加博文
func GetArticle(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if !repository.ArticleExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("Article %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	article := repository.GetArticle(uint(id))
	// 转换为Dto
	c.JSON(http.StatusOK, dtos.ParseArticleEntity(article))
}

func GetArticles(c *gin.Context) {
	articles := repository.GetArticles()

	// 转换为Dto
	articleDtos := make([]dtos.ArticleGetDto, 0, len(articles))
	for _, article := range articles {
		articleDtos = append(articleDtos, *dtos.ParseArticleEntity(&article))
	}

	c.JSON(http.StatusOK, articleDtos)
}

func UpdateArticle(c *gin.Context) {
	// 获得更新id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if !repository.ArticleExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("Article %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}
	article := repository.GetArticle(uint(id))

	// 获得patchJson
	patchJson, getRawDataErr := c.GetRawData()
	if getRawDataErr != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Bind Model Error",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 应用patch
	dto := dtos.ArticleUpdateDtoFromEntity(article)
	utils.ApplyJsonPatch(dto, patchJson)
	dto.ApplyUpdateToEntity(article)

	// 更新数据库
	repository.UpdateArticle(article)

	c.Status(http.StatusNoContent)
}

func DeleteArticle(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if !repository.ArticleExists(uint(id)) {
		c.JSON(http.StatusNotFound, dtos.ErrorDto{
			Message:          fmt.Sprintf("Article %v not found!", id),
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	repository.DeleteArticle(uint(id))
	c.Status(http.StatusNoContent)
}
