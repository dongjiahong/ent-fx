// Package app app包
package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"

	"web/internal/pkg/errs"
)

type Gin struct {
	C *gin.Context
}

// 定义一个全局的翻译器
var trans ut.Translator

func init() {
	trans, _ = ut.New(zh.New()).GetTranslator("zh")
	zhTranslations.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), trans)
}

func ValidateError(err error) (ret error) {
	var retStr string
	if validationErrors, ok := err.(validator.ValidationErrors); !ok {
		return err
	} else {
		for _, e := range validationErrors {
			retStr += e.Translate(trans) + ";"
		}
	}
	return fmt.Errorf("%s", retStr)
}

func errorHandler(err error) (int, string) {
	code := errs.InvalidParams
	msg := errs.MsgFlags[errs.InvalidParams]
	switch v := err.(type) {
	case validator.ValidationErrors:
		msg = ""
		for _, e := range v {
			msg += e.Translate(trans) + ";"
		}
		code = errs.InvalidParams
	case *errs.CodeError:
		msg = v.Error()
		code = v.Code()
	case error:
		msg = v.Error()
	}
	return code, msg
}

func (g *Gin) ResponseJSON(err error, data interface{}) {
	code, msg := errorHandler(err)
	g.C.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func (g *Gin) ResponseJSONSuccess(data interface{}) {
	g.C.JSON(http.StatusOK, gin.H{
		"code": errs.Success,
		"data": data,
		"msg":  errs.MsgFlags[errs.Success],
	})
}
