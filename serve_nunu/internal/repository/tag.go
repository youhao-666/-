package repository

import (
	"context"
	"fmt"
	v1 "server_go/api/v1"
	"server_go/internal/model"

	"github.com/go-sql-driver/mysql"
)

type TagRepository interface {
	CreateTag(ctx context.Context, tag *model.Tag) error
	GetTag(ctx context.Context, tag *model.Tag) error
	QueryAllTag(ctx context.Context) ([]*model.Tag, error)

	DeleteAritcle_Tagid(ctx context.Context, tag_id uint) error
	DeleteTag(ctx context.Context, tagid uint) error
}

func NewTagRepository(
	repository *Repository,
) TagRepository {
	return &tagRepository{
		Repository: repository,
	}
}

type tagRepository struct {
	*Repository
}

func (r *tagRepository) CreateTag(ctx context.Context, tag *model.Tag) error {
	err := r.db.Table("tags").Create(tag).Error
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return v1.ErrTagAlreadyUse
			}
			return err
		}
	}
	return nil
}

func (r *tagRepository) GetTag(ctx context.Context, tag *model.Tag) error {
	err := r.db.Table("tags").Where("id = ?", tag.ID).First(tag).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *tagRepository) QueryAllTag(ctx context.Context) ([]*model.Tag, error) {
	var tags []*model.Tag
	err := r.db.Table("tags").Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *tagRepository) DeleteAritcle_Tagid(ctx context.Context, tag_id uint) error {
	if err := r.db.WithContext(ctx).
		Table("article_tags").
		Where("tag_id = ?", tag_id).
		Unscoped().
		Delete(nil).Error; err != nil { // 传nil表示操作当前表
		return err
	}
	return nil
}

// func (r *tagRepository) DeleteTag(ctx context.Context, tagid uint) error {
// 	if err := r.db.Table("tags").Where("id = ?", tagid).Delete(nil).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

func (r *tagRepository) DeleteTag(ctx context.Context, tagid uint) error {
	// 执行删除操作并获取影响的行数
	// result := r.db.WithContext(ctx).
	// 	Table("tags").
	// 	Where("id = ?", tagid).
	// 	Delete(nil)

	result := r.db.WithContext(ctx).
		Table("tags").
		Where("id = ?", tagid).
		Delete(&model.Tag{}) // 使用 Delete 方法自动处理软删除

	if result.Error != nil {
		return fmt.Errorf("删除标签失败: %w", result.Error)
	}

	// 检查是否实际删除了记录
	if result.RowsAffected == 0 {
		return fmt.Errorf("标签ID %d 不存在", tagid)
	}

	return nil
}
