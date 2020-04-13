package app

import (
	"github.com/fmarinCeiba/bookstore_oauth-api/http"
	"github.com/fmarinCeiba/bookstore_oauth-api/repository/db"
	"github.com/fmarinCeiba/bookstore_oauth-api/repository/rest"
	"github.com/fmarinCeiba/bookstore_oauth-api/services/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(access_token.NewService(rest.NewRepository(), db.NewRepository()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
