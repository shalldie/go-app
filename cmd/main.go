package main

import (
	"fmt"
	"go-app/api"
	"go-app/config"

	"github.com/gin-gonic/gin"
)

func main() {

	env := config.LoadEnv()

	app := gin.Default()

	api.Setup(app)

	addr := fmt.Sprintf(":%d", env.PORT)
	app.Run(addr)
}
