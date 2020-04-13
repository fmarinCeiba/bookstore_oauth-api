package http

import (
	"net/http"

	atDomain "github.com/fmarinCeiba/bookstore_oauth-api/domain/access_token"
	"github.com/fmarinCeiba/bookstore_oauth-api/services/access_token"
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
	var atr atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&atr); err != nil {
		rErr := errors.NewBadRequestError("invalid json")
		c.JSON(rErr.Status, rErr)
		return
	}
	accessToken, err := h.service.Create(atr)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}

func (h *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {
	accessToken, err := h.service.GetByID(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, accessToken)
}
