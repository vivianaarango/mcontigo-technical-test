package service

import (
	"context"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"github.com/google/uuid"
)

func (s *service) Get(
	ctx context.Context,
	UserID uuid.UUID,
	BlogID uuid.UUID,
	Interests []newsletter.Interest,
	Page int,
	MaxPageSize int,
) (*newsletter.Result[*newsletter.Subscription], error) {
	result, err := s.repo.Search(ctx, UserID, BlogID, Interests, MaxPageSize, Page)
	if err != nil {
		return &newsletter.Result[*newsletter.Subscription]{}, err
	}

	response := &newsletter.Result[*newsletter.Subscription]{
		Total: len(result),
		Pages: len(result) / MaxPageSize,
		Page: newsletter.Page[*newsletter.Subscription]{
			Number:   Page,
			Elements: result,
		},
	}

	return response, nil
}
