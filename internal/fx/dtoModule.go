package fx_utils

import (
	"github.com/nihal-ramaswamy/chalk_mvp/internal/dto"
	"go.uber.org/fx"
)

var DtoModule = fx.Module(
	"DtoModule",
	fx.Provide(dto.NewRoom),
)
