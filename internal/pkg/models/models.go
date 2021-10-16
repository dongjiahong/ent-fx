package models

import (
	"context"
	"fmt"
	"time"
	"web/internal/pkg/config"

	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DBHandler interface {
	GetDBHandler() *gorm.DB
	GetDBTxHandler() *gorm.DB
}

type dbHandler struct {
	db   *gorm.DB
	conf *config.DBConfig
}

type Option interface {
	apply(*dbHandler)
}

type optionFunc func(*dbHandler)

func (f optionFunc) apply(d *dbHandler) { f(d) }

func SetConfigOption(conf *config.DBConfig) Option {
	return optionFunc(func(dbh *dbHandler) {
		dbh.conf = conf
	})
}

func NewDBHandler(lc fx.Lifecycle, opts ...Option) DBHandler {
	dbh := dbHandler{}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			for _, opt := range opts {
				opt.apply(&dbh)
			}
			fmt.Println("--------> database start")
			// connect to mysql
			dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				dbh.conf.User, dbh.conf.Password, dbh.conf.Host, dbh.conf.DBName)
			var err error
			dbh.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					TablePrefix:   dbh.conf.TablePrefix,
					SingularTable: true,
				},
				Logger: logger.Default.LogMode(logger.Info),
			})
			if err != nil {
				panic(err)
			}
			sqlDB, err := dbh.db.DB()
			if err != nil {
				panic(err)
			}
			if err := sqlDB.Ping(); err != nil {
				panic(err)
			}
			sqlDB.SetConnMaxIdleTime(10)
			sqlDB.SetMaxOpenConns(50)
			sqlDB.SetConnMaxIdleTime(time.Minute * 5)
			return nil
		},
		OnStop: func(context.Context) error {
			fmt.Println("-------> database stop")
			return nil
		},
	})

	return &dbh
}

func (dbh *dbHandler) GetDBHandler() *gorm.DB {
	return dbh.db
}

func (dbh *dbHandler) GetDBTxHandler() *gorm.DB {
	return dbh.db.Begin()
}
