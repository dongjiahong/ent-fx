package api

import (
	"github.com/gin-gonic/gin"

	"web/internal/pkg/app"
	"web/internal/services/helloserver"
)

// Hello
// @Summary 你好
// @Tags example
// @Success 200 {string} json "{"code": 200}"
// @Success 400 {string} json "{"code": 400}"
// @Router /api/v1/hello [GET]
func Hello(hs helloserver.HelloServer) func(c *gin.Context) {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		appG.ResponseJSONSuccess(hs.SayHello())
	}
}
