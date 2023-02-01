package handler

import (
	"net/http"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// nolint:lll // godoc
// Get godoc
// @Summary      Read subscriptions
// @Tags         subscriptions
// @Router       /subscriptions [get]
// @Param        page	        query  int		 true   "Result page"                                   example(1)
// @Param        maxPageSize	query  int		 true   "Max page size"                                  example(10)
// @Param        userId	        query  string	 false  "User ID"                                        example(1)
// @Param        blogId	        query  string	 false  "Blog ID"                                        example(1)
// @Param        interests	    query  []string  false  "Interests"                                      example(["tech","sports"])
// @Produce      json
// @Success      200  {array}  handler.ResponseDoc
// nolint:gocyclo //error checking branches
func (h *handler) Get(ctx *gin.Context) {
	r := &request{}

	if err := ctx.ShouldBindJSON(r); err != nil {
		return
	}

	userID, err := uuid.Parse(r.Filter.UserID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "user id is not valid")
		return
	}

	blogID, err := uuid.Parse(r.Filter.BlogID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "blog id is not valid")
		return
	}

	var interests []newsletter.Interest
	// TODO get interest to sent

	subscriptions, err := h.svc.Get(ctx, userID, blogID, interests, r.Pagination.Page, r.Pagination.MaxPageSize)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, "newsletter subscriptions is not found")
		return
	}

	ctx.JSON(http.StatusOK, subscriptions)
}
