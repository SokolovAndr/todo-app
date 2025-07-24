package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorRespose(c, http.StatusUnauthorized, "empty Auth Header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorRespose(c, http.StatusUnauthorized, "invalid Auth Header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorRespose(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}
