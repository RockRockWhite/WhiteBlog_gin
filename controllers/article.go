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

var articleRepository *services.ArticleRepository

// InitArticleController 初始化博文Controller
func InitArticleController() {
    articleRepository = services.NewArticleRepository(true)
}

// AddArticle 添加博文
func AddArticle(c *gin.Context) {
    var articleDto dtos.ArticleAddDto

    if err := c.ShouldBind(&articleDto); err != nil {
        c.JSON(http.StatusBadRequest, dtos.ErrorDto{
            Message:          "Bind Model Error",
            DocumentationUrl: viper.GetString("Document.Url"),
        })
        return
    }

    // 获得用户信息
    claims := c.MustGet("claims").(*utils.JwtClaims)

    entity := articleDto.ToEntity(claims.Id)
    articleRepository.AddArticle(entity)

    c.JSON(http.StatusCreated, dtos.ParseArticleEntity(entity))
}

// GetArticle 获得博文
func GetArticle(c *gin.Context) {
    id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

    if !articleRepository.ArticleExists(uint(id)) {
        c.JSON(http.StatusNotFound, dtos.ErrorDto{
            Message:          fmt.Sprintf("Article %v not found!", id),
            DocumentationUrl: viper.GetString("Document.Url"),
        })
        return
    }

    article := articleRepository.GetArticle(uint(id))
    // 转换为Dto
    c.JSON(http.StatusOK, dtos.ParseArticleEntity(article))
}

func GetArticles(c *gin.Context) {
    articles := articleRepository.GetArticles()

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
    if !articleRepository.ArticleExists(uint(id)) {
        c.JSON(http.StatusNotFound, dtos.ErrorDto{
            Message:          fmt.Sprintf("Article %v not found!", id),
            DocumentationUrl: viper.GetString("Document.Url"),
        })
        return
    }
    article := articleRepository.GetArticle(uint(id))

    // 获得用户信息 判断用户是否对该博文具有修改权
    // 修改权: 改博文为用户所有 或 该用户是管理员
    claims := c.MustGet("claims").(*utils.JwtClaims)

    if article.UserId != claims.Id && !claims.IsAdmin {
        c.JSON(http.StatusForbidden, dtos.ErrorDto{
            Message:          "Permission denied for changing this resource!",
            DocumentationUrl: viper.GetString("Document.Url"),
        })
        return
    }

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
    articleRepository.UpdateArticle(article)

    c.Status(http.StatusNoContent)
}

func DeleteArticle(c *gin.Context) {
    id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

    if !articleRepository.ArticleExists(uint(id)) {
        c.JSON(http.StatusNotFound, dtos.ErrorDto{
            Message:          fmt.Sprintf("Article %v not found!", id),
            DocumentationUrl: viper.GetString("Document.Url"),
        })
        return
    }
    article := articleRepository.GetArticle(uint(id))

    // 获得用户信息 判断用户是否对该博文具有修改权
    // 修改权: 改博文为用户所有 或 该用户是管理员
    claims := c.MustGet("claims").(*utils.JwtClaims)

    if article.UserId != claims.Id && !claims.IsAdmin {
        c.JSON(http.StatusForbidden, dtos.ErrorDto{
            Message:          "Permission denied for changing this resource!",
            DocumentationUrl: viper.GetString("Document.Url"),
        })
        return
    }

    articleRepository.DeleteArticle(uint(id))
    c.Status(http.StatusNoContent)
}
