package models

import (
	"context"
	"fmt"

	"go.uber.org/fx"
)

type DB struct{}

func NewModels(lc fx.Lifecycle) DB {
	db := DB{}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("start------------start")
			return nil
		},
		OnStop: func(context.Context) error {
			fmt.Println("end-------------end")
			return nil
		},
	})

	return db
}
