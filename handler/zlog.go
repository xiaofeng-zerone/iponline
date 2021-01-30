package handler

import (
	"go.uber.org/zap"
)

var zlog *zap.SugaredLogger

//InitLogger init log
func InitLogger() {
	logger, _ := zap.NewProduction()
	zlog = logger.Sugar()
}

//UnInitLogger uninit log
func UnInitLogger() {
	zlog.Sync()
}
