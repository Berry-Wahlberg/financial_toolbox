package main

import (
	"financial_toolbox/internal/api"
	"financial_toolbox/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.CORS())

	api.SetupRoutes(r)

	r.Run(":8080")
}
