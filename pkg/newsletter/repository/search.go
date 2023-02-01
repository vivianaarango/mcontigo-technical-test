package repository

import (
	"context"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"github.com/google/uuid"
)

func (r *repository) Search(
	ctx context.Context,
	userID uuid.UUID,
	blogID uuid.UUID,
	interests []newsletter.Interest,
	limit int,
	offset int,
) ([]*newsletter.Subscription, error) {
	var subscriptions []*newsletter.Subscription

	for _, subscription := range r.data {
		if len(subscriptions) == limit {
			return subscriptions, nil
		}

		isTheSameInterest := validateInterestFilter(interests, subscription.Interests)
		if subscription.BlogID == blogID.String() && subscription.UserID == userID.String() && isTheSameInterest {
			subscriptions = append(subscriptions, &newsletter.Subscription{
				UserID:    userID,
				BlogID:    blogID,
				Interests: interests,
			})
		}
	}

	return subscriptions, nil
}

// validateInterestFilter check if the interests sent are the same of the subscription data.
func validateInterestFilter(interestsFilter, interestData []newsletter.Interest) bool {
	count := 0
	for _, q := range interestsFilter {
		for _, i := range interestData {
			if q == i {
				count++
			}
		}
	}

	return count == len(interestsFilter)
}
