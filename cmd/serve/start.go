package serve

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"web/internal/pkg/config"
	"web/internal/pkg/models"
	"web/internal/pkg/zlog"
	"web/internal/services/helloserver"
	"web/internal/webserve"
)

var startCMD = &cobra.Command{
	Use:   "start",
	Short: "start web serve",
	Run: func(cmd *cobra.Command, args []string) {
		start(context.Background())
	},
}

func startWeb(conf config.Config, l zlog.Logger, hs helloserver.HelloServer) webserve.WebServe {
	return webserve.NewWebServe(
		hs,
		webserve.SetConfigOption(conf.GetWebConfig()),
		webserve.SetLoggerOption(l),
	)
}

func newHelloServer(l zlog.Logger, dbh models.DBHandler) helloserver.HelloServer {
	return helloserver.NewHelloServer(
		helloserver.SetLoggerOption(l),
		helloserver.SetDBOption(dbh),
	)
}

func newDBHandler(lc fx.Lifecycle, conf config.Config) models.DBHandler {
	return models.NewDBHandler(
		lc,
		models.SetConfigOption(conf.GetDBConfig()),
	)
}

func register(lc fx.Lifecycle, s webserve.WebServe) {
	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				//
				s.ListenAndServe()
				return nil
			},
			OnStop: func(context.Context) error {
				//
				s.Shutdown()
				return nil
			},
		},
	)
}

func start(ctx context.Context) {
	app := fx.New(
		// fx.NopLogger, // 屏蔽日志
		fx.Provide(
			config.NewConfig,
			zlog.NewLogger,
			newDBHandler,
			newHelloServer,
			startWeb,
		),
		fx.Invoke(register),
	)
	if err := app.Start(ctx); err != nil {
		panic(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit
	log.Printf("shuting down server... signal: %s\n", s)

	// graceful shutdown time is 10 seconds
	ctx, cf := context.WithTimeout(context.Background(), time.Second*10)
	defer cf()

	if err := app.Stop(ctx); err != nil {
		log.Println("fatal error occurred, err: ", err)
	}
}
