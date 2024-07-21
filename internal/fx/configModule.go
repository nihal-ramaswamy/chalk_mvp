package fx_utils

import (
	db_config "github.com/nihal-ramaswamy/chalk_mvp/internal/config/db"
	serverconfig "github.com/nihal-ramaswamy/chalk_mvp/internal/config/server"
	"go.uber.org/fx"
)

var ConfigModule = fx.Module(
	"Config",
	fx.Provide(serverconfig.Default),
	fx.Provide(db_config.GetPsqlInfoDefault),
	fx.Provide(
		fx.Annotate(
			db_config.DefaultRedisAuthConfig,
			fx.ResultTags(`name:"auth_rdb_config"`),
		),
	),
	fx.Provide(
		fx.Annotate(
			db_config.DefaultRedisWebsocketConfig,
			fx.ResultTags(`name:"ws_rdb_config"`),
		),
	),
)
