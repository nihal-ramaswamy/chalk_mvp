package utils

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func HandleErrorAndAbortWithError(ctx *gin.Context, err error, log *zap.Logger, code int) {
	if nil != err {
		ctx.Error(err)
		log.Error(err.Error())
		_ = ctx.AbortWithError(code, err)
	}
}
