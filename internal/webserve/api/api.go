package api

import (
	"web/internal/pkg/app"

	"github.com/gin-gonic/gin"
)

// Hello
// @Summary 你好
// @Tags example
// @Success 200 {string} json "{"code": 200}"
// @Router /api/v1/hello [GET]
func Hello(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.ResponseJSONSuccess("world")
}
