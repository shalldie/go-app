package main

import (
	"fmt"
	"go-app/configs"
	"go-app/pkg"

	"github.com/gin-gonic/gin"
)

func main() {

	env := configs.LoadEnv()

	app := gin.Default()

	pkg.Setup(app)

	addr := fmt.Sprintf(":%d", env.PORT)
	app.Run(addr)
}
