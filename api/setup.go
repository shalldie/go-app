package api

import (
	"go-app/api/tianqi"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {

	app.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "All apis start with `/api`")

	})

	router := app.Group("/api")
	tianqi.Setup(router)
}
