package handler

import (
	"net/http"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"github.com/gin-gonic/gin"
)

func (h *handler) Create(ctx *gin.Context) {
	r := &newsletter.Subscription{}

	if err := ctx.ShouldBindJSON(r); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "error getting subscription data")
		return
	}

	err := h.svc.Create(ctx, r.UserID, r.BlogID, r.Interests)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "error creating subscription")
		return
	}

	ctx.JSON(http.StatusOK, "subscription created")
}
