package service

import (
	"context"
	"fmt"
	v1 "server_go/api/v1"
	"server_go/internal/model"
	"server_go/internal/repository"
)

type ArticleService interface {
	GetArticle(ctx context.Context, id int64) (*model.Article, error)
	CreateArticle(ctx context.Context, articleWithTags *v1.ArticleWithTags) error
	QueryArticleByTag(ctx context.Context, tagID uint) ([]*v1.ArticleWithTags, error)
	QueryAllArticle(ctx context.Context) ([]*v1.ArticleWithTags, error)
	QueryArticleByUserID(ctx context.Context, userID uint) ([]*v1.ArticleWithTags, error)
	DeleteArticle(ctx context.Context, Articleid uint) error
}

func NewArticleService(
	service *Service,
	articleRepository repository.ArticleRepository,
) ArticleService {
	return &articleService{
		Service:           service,
		articleRepository: articleRepository,
	}
}

type articleService struct {
	*Service
	articleRepository repository.ArticleRepository
}

func (s *articleService) GetArticle(ctx context.Context, id int64) (*model.Article, error) {
	return s.articleRepository.GetArticle(ctx, id)
}

// func (s *articleService) CreateArticle(ctx context.Context, articleWithTags *v1.ArticleWithTags) error {
// 	err := s.tm.Transaction(ctx, func(ctx context.Context) error {
// 		var article = &model.Article{
// 			Title:   articleWithTags.Title,
// 			Content: articleWithTags.Content,
// 			UserId:  articleWithTags.UserId,
// 		}
// 		var tags = make([]model.Tag, len(articleWithTags.Tags))
// 		for i, tag := range articleWithTags.Tags {
// 			tags[i] = model.Tag{
// 				TagName: tag,
// 			}
// 		}
// 		articleId, err := s.articleRepository.CreateArticle(ctx, article)
// 		if err != nil {
// 			return err
// 		}
// 		var tagId uint
// 		for _, tag := range tags {
// 			if tagId, err = s.articleRepository.CheckTag(ctx, &tag); err != nil {
// 				tagId, err = s.articleRepository.CreateTag(ctx, &tag)
// 				if err != nil {
// 					return err
// 				}
// 				err := s.articleRepository.CreateArticleTag(ctx, articleId, tagId)
// 				if err != nil {
// 					return err
// 				}

// 			} else {
// 				err := s.articleRepository.CreateArticleTag(ctx, articleId, tagId)
// 				if err != nil {
// 					return err
// 				}
// 			}
// 		}

// 		return nil
// 	})
// 	return err
// }

func (s *articleService) CreateArticle(ctx context.Context, articleWithTags *v1.ArticleWithTags) error {
	return s.tm.Transaction(ctx, func(ctx context.Context) error {

		//  处理标签
		for _, tagName := range articleWithTags.Tags {

			tag := &model.Tag{TagName: tagName}
			_, err := s.getOrCreateTag(ctx, tag)
			if err != nil {
				return err
			}
		}

		// 1. 创建文章（只执行一次）
		article := &model.Article{
			Title:   articleWithTags.Title,
			Content: articleWithTags.Content,
			UserId:  articleWithTags.UserId,
		}

		articleID, err := s.articleRepository.CreateArticle(ctx, article)
		if err != nil {
			return err
		}

		//  处理标签
		for _, tagName := range articleWithTags.Tags {

			tag := &model.Tag{TagName: tagName}
			tagID, err := s.getOrCreateTag(ctx, tag)
			if err != nil {
				return err
			}

			if err := s.safeCreateArticleTag(ctx, articleID, tagID); err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *articleService) safeCreateArticleTag(ctx context.Context, articleID, tagID uint) error {
	// 使用防重插入
	err := s.articleRepository.CreateArticleTag(ctx, articleID, tagID)
	if err == nil {
		return nil
	}
	return err
}

func (s *articleService) getOrCreateTag(ctx context.Context, tag *model.Tag) (uint, error) {
	tagID, err := s.articleRepository.CheckTag(ctx, tag)
	if err == nil {
		return tagID, nil
	} else {
		return 0, v1.ErrTagNntFound
		//return s.articleRepository.CreateTag(ctx, tag)
	}

}

func (s *articleService) QueryArticleByTag(ctx context.Context, tagID uint) ([]*v1.ArticleWithTags, error) {
	//
	articleIds, err := s.articleRepository.QueryArticleIdByTag(ctx, tagID)
	if err != nil {
		return nil, err
	}
	var articleWithTagsList = make([]*v1.ArticleWithTags, len(articleIds))
	for i, articleId := range articleIds {
		article, err := s.QueryOneArticle(ctx, articleId)
		if err != nil {
			return nil, err
		}
		articleWithTagsList[i] = article
	}
	return articleWithTagsList, nil
}

func (s *articleService) QueryOneArticle(ctx context.Context, id uint) (*v1.ArticleWithTags, error) {

	Tags, err := s.articleRepository.QueryTagByArticleID(ctx, id)
	if err != nil {
		return nil, err
	}
	artical, err := s.articleRepository.QueryArticle(ctx, id)
	if err != nil {
		return nil, err
	}
	var articleWithTags = v1.ArticleWithTags{
		ID:      artical.ID,
		Title:   artical.Title,
		Content: artical.Content,
		UserId:  artical.UserId,
		Tags:    make([]string, len(Tags)),
	}
	for i, tag := range Tags {
		articleWithTags.Tags[i] = tag.TagName
	}
	return &articleWithTags, nil
}

func (s *articleService) QueryAllArticle(ctx context.Context) ([]*v1.ArticleWithTags, error) {
	ArticleIds, err := s.articleRepository.QueryAllArticleId(ctx)
	if err != nil {
		return nil, err
	}
	var articleWithTagsList = make([]*v1.ArticleWithTags, len(ArticleIds))
	for i, ArticleId := range ArticleIds {
		article, err := s.QueryOneArticle(ctx, ArticleId)
		if err != nil {
			return nil, err
		}
		articleWithTagsList[i] = article
	}
	return articleWithTagsList, nil
}

func (s *articleService) QueryArticleByUserID(ctx context.Context, userID uint) ([]*v1.ArticleWithTags, error) {
	//检查用户是否存在
	err := s.articleRepository.CheckUserid(ctx, userID)
	if err != nil {
		return nil, err
	}
	ArticleIds, err := s.articleRepository.QueryArticleIdByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	fmt.Println(userID)
	fmt.Println(ArticleIds)
	var articleWithTagsList = make([]*v1.ArticleWithTags, len(ArticleIds))
	for i, ArticleId := range ArticleIds {
		article, err := s.QueryOneArticle(ctx, ArticleId)
		if err != nil {
			return nil, err
		}
		articleWithTagsList[i] = article
	}
	return articleWithTagsList, nil
}

func (s *articleService) DeleteArticle(ctx context.Context, Articleid uint) error {
	return s.tm.Transaction(ctx, func(ctx context.Context) error {
		if err := s.articleRepository.DeleteAritcle_Tagid(ctx, Articleid); err != nil {
			return err
		}
		if err := s.articleRepository.DeleteArticle(ctx, Articleid); err != nil {
			return err
		}
		return nil
	})
}
