package helloserver

import (
	"fmt"
	"web/internal/pkg/models"
	"web/internal/pkg/models/sealed"
	"web/internal/pkg/zlog"
)

type HelloServer interface {
	SayHello() string
}

type helloServer struct {
	l   zlog.Logger
	dbh models.DBHandler
}

type Option interface {
	apply(*helloServer)
}

type optionFunc func(*helloServer)

func (f optionFunc) apply(h *helloServer) { f(h) }

func SetLoggerOption(l zlog.Logger) Option {
	return optionFunc(func(hs *helloServer) {
		hs.l = l
	})
}

func SetDBOption(dbh models.DBHandler) Option {
	return optionFunc(func(hs *helloServer) {
		hs.dbh = dbh
	})
}

func NewHelloServer(opts ...Option) HelloServer {
	instance := helloServer{}
	for _, opt := range opts {
		opt.apply(&instance)
	}

	return &instance
}

func (hs *helloServer) SayHello() string {
	hs.l.Debug("========> db: ", hs.dbh.GetDBHandler())
	s, err := sealed.GetSealed(hs.dbh.GetDBHandler(), "s-t1009-100")
	if err != nil {
		hs.l.Error("get sealed err: ", err)
	}
	return fmt.Sprintf("name is %s, ticket: %s", s.Name, s.Ticket)
}
