package fx_utils

import (
	"context"
	"database/sql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	serverconfig "github.com/nihal-ramaswamy/chalk_mvp/internal/config/server"
	middleware_log "github.com/nihal-ramaswamy/chalk_mvp/internal/middleware/log"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/routes"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func newServerEngine(
	lc fx.Lifecycle,
	rdb_auth *redis.Client,
	rdb_ws *redis.Client,
	config *serverconfig.Config,
	log *zap.Logger,
	db *sql.DB,
	ctx context.Context,
) *gin.Engine {
	gin.SetMode(config.GinMode)

	server := gin.Default()
	server.Use(cors.New(config.Cors))
	server.Use(middleware_log.DefaultStructuredLogger(log))
	server.Use(gin.Recovery())

	routes.NewRoutes(server, log, db, rdb_auth, ctx, rdb_ws)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Starting server on port", zap.String("port", config.Port))

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stopping server")
			defer func() {
				err := log.Sync()
				if nil != err {
					log.Error(err.Error())
				}
			}()

			return nil
		},
	})

	return server
}

var serverModule = fx.Module(
	"serverModule",
	fx.Provide(
		fx.Annotate(
			newServerEngine,
			fx.ParamTags(``, `name:"auth_rdb"`, `name:"ws_rdb"`),
		),
	),
)
