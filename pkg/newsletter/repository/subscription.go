package repository

import "git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"

type subscriptionDBModel struct {
	UserID    string
	BlogID    string
	Interests []newsletter.Interest
}
