package fx_utils

import (
	websockets_impl "github.com/nihal-ramaswamy/chalk_mvp/internal/websockets"
	"go.uber.org/fx"
)

var DtoModule = fx.Module(
	"DtoModule",
	fx.Provide(websockets_impl.NewRoom),
)
