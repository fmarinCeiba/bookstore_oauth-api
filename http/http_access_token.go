package http

import (
	"net/http"

	"github.com/fmarinCeiba/bookstore_oauth-api/domain/access_token"
	"github.com/fmarinCeiba/bookstore_oauth-api/utils/errors"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
	UpdateExpirationTime(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(s access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: s,
	}
}

func (h *accessTokenHandler) GetByID(c *gin.Context) {
	accessToken, err := h.service.GetByID(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		rErr := errors.NewBadRequestError("invalid json")
		c.JSON(rErr.Status, rErr)
		return
	}
	if err := h.service.Create(at); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, at)
}

func (h *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {
	accessToken, err := h.service.GetByID(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, accessToken)
}
