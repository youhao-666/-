package repository

import (
	"context"
	"fmt"
	v1 "server_go/api/v1"
	"server_go/internal/model"
)

type ArticleRepository interface {
	GetArticle(ctx context.Context, id int64) (*model.Article, error)
	CreateArticle(ctx context.Context, article *model.Article) (uint, error)
	CheckTag(ctx context.Context, tag *model.Tag) (uint, error)
	CreateTag(ctx context.Context, tag *model.Tag) (uint, error)
	CreateArticleTag(ctx context.Context, articleID uint, tagID uint) error
	QueryTagByArticleID(ctx context.Context, articleID uint) ([]*model.Tag, error)
	QueryArticle(ctx context.Context, id uint) (*model.Article, error)
	QueryArticleIdByTag(ctx context.Context, tagID uint) ([]uint, error)
	QueryAllArticleId(ctx context.Context) ([]uint, error)
	QueryArticleIdByUserID(ctx context.Context, userID uint) ([]uint, error)
	CheckUserid(ctx context.Context, userID uint) error
	DeleteAritcle_Tagid(ctx context.Context, articleID uint) error
	DeleteArticle(ctx context.Context, id uint) error
}

func NewArticleRepository(
	repository *Repository,
) ArticleRepository {
	return &articleRepository{
		Repository: repository,
	}
}

type articleRepository struct {
	*Repository
}

func (r *articleRepository) GetArticle(ctx context.Context, id int64) (*model.Article, error) {
	var article model.Article

	return &article, nil
}

func (r *articleRepository) CreateArticle(ctx context.Context, article *model.Article) (uint, error) {
	// 只需要Create操作
	if err := r.db.Table("articles").WithContext(ctx).Create(article).Error; err != nil {
		return 0, err
	}
	// Create成功后，article的ID会被自动填充
	return article.ID, nil
}

func (r *articleRepository) CheckTag(ctx context.Context, tag *model.Tag) (uint, error) {
	t := new(model.Tag)
	if err := r.db.Table("tags").Where("tag_name = ?", tag.TagName).First(t).Error; err != nil {
		return 0, err
	}
	return t.ID, nil
}

func (r *articleRepository) CreateTag(ctx context.Context, tag *model.Tag) (uint, error) {

	t := new(model.Tag)
	if err := r.db.Table("tags").Create(tag).First(t).Error; err != nil {
		return 0, err
	}
	return t.ID, nil
}

func (r *articleRepository) CreateArticleTag(ctx context.Context, articleID uint, tagID uint) error {
	articleTag := model.ArticleTag{
		ArticleID: articleID,
		TagID:     tagID,
	}
	if err := r.db.Table("article_tags").Create(articleTag).Error; err != nil {
		return err
	}
	return nil
}

func (r *articleRepository) QueryTagByArticleID(ctx context.Context, articleID uint) ([]*model.Tag, error) {
	var tags []*model.Tag
	err := r.db.Table("tags").
		Joins("JOIN article_tags ON article_tags.tag_id = tags.id").
		Where("article_tags.article_id = ?", articleID).
		Find(&tags).Error
	return tags, err
}

func (r *articleRepository) QueryArticle(ctx context.Context, id uint) (*model.Article, error) {
	art := new(model.Article)
	if err := r.db.Table("articles").WithContext(ctx).First(art, id).Error; err != nil {
		return nil, err
	}
	return art, nil
}

func (r *articleRepository) QueryArticleIdByTag(ctx context.Context, tagID uint) ([]uint, error) {
	var articleIds []uint

	// 使用WithContext添加上下文，并查询article_tags表中给定tagID的所有articleID
	if err := r.db.Table("article_tags").
		Where("tag_id = ?", tagID).
		Pluck("article_id", &articleIds).
		Error; err != nil {
		return nil, err
	}
	return articleIds, nil
}

func (r *articleRepository) QueryAllArticleId(ctx context.Context) ([]uint, error) {
	var articleIds []uint

	if err := r.db.Table("articles").
		Where("deleted_at IS NULL").
		Pluck("id", &articleIds).
		Error; err != nil {
		return nil, err
	}
	return articleIds, nil
}

func (r *articleRepository) QueryArticleIdByUserID(ctx context.Context, userID uint) ([]uint, error) {
	// 查询给定userID的所有文章ID
	var articleIds []uint
	if err := r.db.Table("articles").
		Where("user_id = ? AND deleted_at IS NULL", userID).
		Pluck("id", &articleIds).
		Error; err != nil {
		return nil, err
	}
	return articleIds, nil
}

func (r *articleRepository) CheckUserid(ctx context.Context, userID uint) error {
	var count int64
	if err := r.db.Table("user").Where("id = ?", userID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return v1.ErrUserNotFound
	}
	return nil
}

func (r *articleRepository) DeleteAritcle_Tagid(ctx context.Context, articleID uint) error {
	if err := r.db.Table("article_tags").Where("article_id = ?", articleID).Delete(model.ArticleTag{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *articleRepository) DeleteArticle(ctx context.Context, id uint) error {
	// if err := r.db.Table("articles").WithContext(ctx).Delete(model.Article{}, id).Error; err != nil {
	// 	return err
	// }
	// return nil

	// 执行删除操作并获取影响的行数
	result := r.db.WithContext(ctx).
		Table("articles"). // 使用模型而不是直接指定表名
		Where("id = ?", id).
		Delete(&model.Article{}) // 传递模型实例

	if result.Error != nil {
		return fmt.Errorf("删除文章失败: %w", result.Error)
	}

	// 检查是否实际更新了记录（软删除是更新操作）
	if result.RowsAffected == 0 {
		return fmt.Errorf("文章ID %d 不存在", id)
	}

	return nil
}
