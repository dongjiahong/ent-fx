package webserve

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "web/docs"
	"web/internal/pkg/config"
	"web/internal/pkg/zlog"
	"web/internal/webserve/api"
)

type WebServe interface {
	ListenAndServe()
	Shutdown()
}

type webServe struct {
	server *http.Server
	l      zlog.Logger
	conf   *config.WebConfig
}

type Option interface {
	apply(*webServe)
}

type optionFunc func(*webServe)

func (f optionFunc) apply(ws *webServe) {
	f(ws)
}

// 设置web的端口
func SetConfigOption(conf *config.WebConfig) Option {
	return optionFunc(func(ws *webServe) {
		ws.conf = conf
	})
}

func SetLoggerOption(l zlog.Logger) Option {
	return optionFunc(func(ws *webServe) {
		ws.l = l
	})
}

func NewWebServe(opts ...Option) WebServe {
	instance := &webServe{}
	for _, opt := range opts {
		opt.apply(instance)
	}
	if instance.server == nil {
		instance.server = newServer()
	}
	if instance.conf.EndPoint == 0 {
		// 默认为2354
		instance.conf.EndPoint = 2354
	}
	return instance
}

func newServer() *http.Server {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := r.Group("api/v1")
	apiV1.GET("/hello", api.Hello)

	server := &http.Server{
		Handler: r,
		// ReadTimeout: ,
		// WriteTimeout: ,
	}

	return server
}

func (ws *webServe) ListenAndServe() {
	ws.server.Addr = fmt.Sprintf("0.0.0.0:%d", ws.conf.EndPoint)

	ws.l.Infof("start http server listening %d\n", ws.conf.EndPoint)
	if err := ws.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

func (ws *webServe) Shutdown() {
	if err := ws.server.Shutdown(context.Background()); err != nil {
		ws.l.Warn("web server forced to shutdown")
	} else {
		ws.l.Info("web server gracefully existed")
	}
}
