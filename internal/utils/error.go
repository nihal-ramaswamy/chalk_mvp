package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func HandleErrorAndAbortWithError(ctx *gin.Context, err error, log *zap.Logger) {
	if nil != err {
		log.Error(err.Error())
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
	}
}
