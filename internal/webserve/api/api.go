package api

import (
	"web/internal/pkg/app"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.ResponseJSONSuccess("world")
}
