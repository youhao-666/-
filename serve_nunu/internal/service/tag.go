package service

import (
	"context"
	v1 "server_go/api/v1"
	"server_go/internal/model"
	"server_go/internal/repository"
)

type TagService interface {
	CreateTag(ctx context.Context, tag *v1.TagCreate) error
	QueryAllTag(ctx context.Context) ([]*model.Tag, error)
	DeleteTag(ctx context.Context, tagid uint) error
}

func NewTagService(
	service *Service,
	tagRepository repository.TagRepository,
) TagService {
	return &tagService{
		Service:       service,
		tagRepository: tagRepository,
	}
}

type tagService struct {
	*Service
	tagRepository repository.TagRepository
}

func (s *tagService) CreateTag(ctx context.Context, tag *v1.TagCreate) error {
	if err := s.tagRepository.GetTag(ctx, &model.Tag{TagName: tag.TagName}); err == nil {
		return v1.ErrTagDefinded
	}
	return s.tagRepository.CreateTag(ctx, &model.Tag{TagName: tag.TagName})
}

func (s *tagService) QueryAllTag(ctx context.Context) ([]*model.Tag, error) {
	var tags []*model.Tag
	tags, err := s.tagRepository.QueryAllTag(ctx)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (s *tagService) DeleteTag(ctx context.Context, tagid uint) error {
	return s.tm.Transaction(ctx, func(ctx context.Context) error {
		if err := s.tagRepository.DeleteAritcle_Tagid(ctx, tagid); err != nil {
			return err
		}
		if err := s.tagRepository.DeleteTag(ctx, tagid); err != nil {
			return err
		}
		return nil
	})
}
