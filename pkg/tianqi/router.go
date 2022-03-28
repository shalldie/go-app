package tianqi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const UID = "U785B76FC9"

func Setup(router *gin.RouterGroup) {
	router.GET("/tianqi/:city", func(ctx *gin.Context) {
		city := ctx.Param("city")

		tuples := getTianqiTuplesByCity(city)

		ctx.JSON(http.StatusOK, gin.H{
			"city":        tuples[0],
			"weather":     tuples[1],
			"temperature": tuples[2],
		})
	})
}
