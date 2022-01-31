package services

import (
	"fmt"
	"gin/entities"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

// NewArticleRepository 创建新博客Repository
func NewArticleRepository(autoMigrate bool) *ArticleRepository {
	Host := viper.GetString("DataBase.Host")
	Port := viper.GetString("DataBase.Port")
	Username := viper.GetString("DataBase.Username")
	Password := viper.GetString("DataBase.Password")
	DBName := viper.GetString("DataBase.DBName")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Username, Password, Host, Port, DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("Fatal error open database:%s %s \n", dsn, err))
	}

	// 完成Article迁移
	if autoMigrate {
		if err := db.AutoMigrate(&entities.Article{}); err != nil {
			panic(fmt.Errorf("Fatal migrate database %s : %s \n", "Article", err))
		}

		if err := db.AutoMigrate(&entities.Star{}); err != nil {
			panic(fmt.Errorf("Fatal migrate database %s : %s \n", "Star", err))
		}

		if err := db.AutoMigrate(&entities.Tag{}); err != nil {
			panic(fmt.Errorf("Fatal migrate database %s : %s \n", "Tag", err))
		}

		if err := db.AutoMigrate(&entities.Comment{}); err != nil {
			panic(fmt.Errorf("Fatal migrate database %s : %s \n", "Comment", err))
		}
	}

	repository := ArticleRepository{db}
	return &repository
}

// GetArticle  从id获得博文
func (repository *ArticleRepository) GetArticle(id uint) (*entities.Article, error) {
	if !repository.ArticleExists(id) {
		panic(fmt.Errorf("article id %v not exists", id))
	}
	var article entities.Article
	result := repository.db.First(&article, id)

	// 获得子成员
	article.Stars, _ = repository.GetStars(id)
	article.Tags, _ = repository.GetTags(id)
	article.Comments, _ = repository.GetComments(id)

	return &article, result.Error
}

func (repository *ArticleRepository) GetArticles() ([]entities.Article, error) {
	var articles []entities.Article
	result := repository.db.Find(&articles)

	// 获得各个子成员
	for i, _ := range articles {
		articles[i].Stars, _ = repository.GetStars(articles[i].ID)
		articles[i].Tags, _ = repository.GetTags(articles[i].ID)
		articles[i].Comments, _ = repository.GetComments(articles[i].ID)
	}

	return articles, result.Error
}

// AddArticle 创建博文
func (repository *ArticleRepository) AddArticle(article *entities.Article) (uint, error) {
	result := repository.db.Create(article)

	// 添加Tags
	for i, _ := range article.Tags {
		article.Tags[i].ArticleId = article.ID
		if _, err := repository.AddTag(&article.Tags[i]); err != nil {
			panic(fmt.Errorf("fatal add tag %+v : %s", article.Tags[i], err))
		}
	}

	return article.ID, result.Error
}

// UpdateArticle 更新博文
func (repository *ArticleRepository) UpdateArticle(article *entities.Article) error {
	result := repository.db.Save(&article)

	return result.Error
}

// DeleteArticle 删除博文
func (repository *ArticleRepository) DeleteArticle(id uint) error {
	if !repository.ArticleExists(id) {
		panic(fmt.Errorf("article id %v not exists", id))
	}
	result := repository.db.Delete(&entities.Article{}, id)

	return result.Error
}

// CommitChanges 提交博文数据库事务修改
func (repository *ArticleRepository) CommitChanges() {
	repository.db.Commit()
}

// ArticleExists 判断该id是否存在
func (repository *ArticleRepository) ArticleExists(id uint) bool {
	var article entities.Article
	result := repository.db.First(&article, id)

	return result.RowsAffected >= 1
}

// GetStar 获得单个点赞
func (repository *ArticleRepository) GetStar(id uint) (*entities.Star, error) {
	if !repository.StarExists(id) {
		panic(fmt.Errorf("star id %v not exists", id))
	}

	var star entities.Star
	result := repository.db.First(&star, id)

	return &star, result.Error
}

// GetStars 获得博文的所有点赞
func (repository *ArticleRepository) GetStars(articleId uint) ([]entities.Star, error) {
	if !repository.ArticleExists(articleId) {
		panic(fmt.Errorf("article id %v not exists", articleId))
	}

	var stars []entities.Star
	result := repository.db.Where(&entities.Star{ArticleId: articleId}).Find(&stars)

	return stars, result.Error
}

// AddStar 添加一条点赞记录
func (repository *ArticleRepository) AddStar(star *entities.Star) (uint, error) {
	if !repository.ArticleExists(star.ArticleId) {
		panic(fmt.Errorf("article id %v not exists", star.ArticleId))
	}

	result := repository.db.Create(star)

	return star.ID, result.Error
}

// UpdateStar 更新一条点赞记录
func (repository *ArticleRepository) UpdateStar(star *entities.Star) error {
	if !repository.ArticleExists(star.ArticleId) {
		panic(fmt.Errorf("article id %v not exists", star.ArticleId))
	}

	result := repository.db.Save(star)

	return result.Error
}

// DeleteStar 删除一条点赞记录
func (repository *ArticleRepository) DeleteStar(id uint) error {
	if !repository.StarExists(id) {
		panic(fmt.Errorf("star id %v not exists", id))
	}

	var star entities.Star
	result := repository.db.Delete(&star, id)

	return result.Error
}

// StarExists 判断该点赞是否存在
func (repository *ArticleRepository) StarExists(id uint) bool {
	var star entities.Star
	result := repository.db.First(&star, id)

	return result.RowsAffected >= 1
}

// GetTag 获得单个F的Tag
func (repository *ArticleRepository) GetTag(id uint) (*entities.Tag, error) {
	if !repository.TagExists(id) {
		panic(fmt.Errorf("tag id %v not exists", id))
	}

	var tag entities.Tag
	result := repository.db.First(&tag, id)

	return &tag, result.Error
}

// GetTags 获得博客的所有Tag
func (repository *ArticleRepository) GetTags(articleId uint) ([]entities.Tag, error) {
	if !repository.ArticleExists(articleId) {
		panic(fmt.Errorf("article id %v not exists", articleId))
	}

	var tags []entities.Tag
	result := repository.db.Where(&entities.Tag{ArticleId: articleId}).Find(&tags)

	return tags, result.Error
}

// AddTag 增加一条Tag
func (repository *ArticleRepository) AddTag(tag *entities.Tag) (uint, error) {
	if !repository.ArticleExists(tag.ArticleId) {
		panic(fmt.Errorf("article id %v not exists", tag.ArticleId))
	}

	result := repository.db.Create(tag)

	return tag.ID, result.Error
}

// UpdateTag 更新Tag
func (repository *ArticleRepository) UpdateTag(tag *entities.Tag) error {
	if !repository.ArticleExists(tag.ArticleId) {
		panic(fmt.Errorf("article id %v not exists", tag.ArticleId))
	}

	result := repository.db.Save(tag)

	return result.Error
}

// DeleteTag 删除Tag
func (repository *ArticleRepository) DeleteTag(id uint) error {
	if !repository.TagExists(id) {
		panic(fmt.Errorf("tag id %v not exists", id))
	}

	var tag entities.Tag
	result := repository.db.Delete(&tag, id)

	return result.Error
}

// TagExists Tag是否存在
func (repository *ArticleRepository) TagExists(id uint) bool {
	var tag entities.Tag
	result := repository.db.First(&tag, id)

	return result.RowsAffected >= 1
}

// GetComment 获得评论
func (repository *ArticleRepository) GetComment(id uint) (*entities.Comment, error) {
	if !repository.CommentExists(id) {
		panic(fmt.Errorf("comment id %v not exists", id))
	}

	var comment entities.Comment
	result := repository.db.First(&comment, id)

	return &comment, result.Error
}

// GetComments 获得评论
func (repository *ArticleRepository) GetComments(articleId uint) ([]entities.Comment, error) {
	if !repository.ArticleExists(articleId) {
		panic(fmt.Errorf("article id %v not exists", articleId))
	}

	var articles []entities.Comment
	result := repository.db.Where(&entities.Comment{ArticleId: articleId}).Find(&articles)

	return articles, result.Error
}

// AddComment 添加一条评论
func (repository *ArticleRepository) AddComment(comment *entities.Comment) (uint, error) {
	if !repository.ArticleExists(comment.ArticleId) {
		panic(fmt.Errorf("article id %v not exists", comment.ArticleId))
	}

	result := repository.db.Create(comment)

	return comment.ID, result.Error
}

// UpdateComment 更新一条评论
func (repository *ArticleRepository) UpdateComment(comment *entities.Comment) error {
	if !repository.ArticleExists(comment.ArticleId) {
		panic(fmt.Errorf("article id %v not exists", comment.ArticleId))
	}

	result := repository.db.Save(comment)

	return result.Error
}

// DeleteComment 删除一条评论
func (repository *ArticleRepository) DeleteComment(id uint) error {
	if !repository.CommentExists(id) {
		panic(fmt.Errorf("comment id %v not exists", id))
	}
	var comment entities.Comment
	result := repository.db.Delete(&comment, id)
	return result.Error
}

// CommentExists 判断评论否存在
func (repository *ArticleRepository) CommentExists(id uint) bool {
	var comment entities.Comment
	result := repository.db.First(&comment, id)

	return result.RowsAffected >= 1
}
