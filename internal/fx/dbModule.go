package fx_utils

import (
	"context"

	"github.com/nihal-ramaswamy/chalk_mvp/internal/db"
	"go.uber.org/fx"
)

var dbModule = fx.Module(
	"DatabaseServices",
	fx.Provide(func() context.Context {
		return context.Background()
	}),

	fx.Provide(db.GetPostgresDbInstanceWithConfig),
	fx.Provide(
		fx.Annotate(
			db.GetRedisDbInstanceWithConfig,
			fx.ParamTags(`name:"auth_rdb_config"`),
			fx.ResultTags(`name:"auth_rdb"`),
		),
	),
	fx.Provide(
		fx.Annotate(
			db.GetRedisDbInstanceWithConfig,
			fx.ParamTags(`name:"ws_rdb_config"`),
			fx.ResultTags(`name:"ws_rdb"`),
		),
	),
)
