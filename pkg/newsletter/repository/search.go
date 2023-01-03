package repository

import (
	"context"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	uuid "github.com/google/uuid"
)

func (r *repository) Search(
	ctx context.Context,
	userID uuid.UUID,
	blogID uuid.UUID,
	interests []newsletter.Interest,
	limit int,
	offset int,
) ([]*newsletter.Subscription, error) {
	panic("implement me")
}
