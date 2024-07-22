package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	serverconfig "github.com/nihal-ramaswamy/chalk_mvp/internal/config/server"
	fx_utils "github.com/nihal-ramaswamy/chalk_mvp/internal/fx"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/utils"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(utils.NewProduction),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),

		fx_utils.ConfigModule,
		fx_utils.MicroServicesModule,
		fx_utils.DtoModule,

		fx.Invoke(Invoke),
	).Run()
}

func Invoke(server *gin.Engine, config *serverconfig.Config, log *zap.Logger) {
	err := server.Run(config.Port)
	if nil != err {
		log.Error(err.Error())
	}
}
