package repository

import (
	"context"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"github.com/google/uuid"
)

func (r *repository) Create(
	ctx context.Context,
	userID uuid.UUID,
	blogID uuid.UUID,
	interests []newsletter.Interest,
) (err error) {
	r.data = append(r.data, &subscriptionDBModel{
		UserID:    userID.String(),
		BlogID:    blogID.String(),
		Interests: interests,
	})

	return err
}
