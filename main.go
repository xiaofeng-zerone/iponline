package main


import (
	"go.uber.org/zap"
)

var zlog *zap.SugaredLogger

func InitLogger() {
	logger, _ := zap.NewProduction()
	zlog = logger.Sugar()
}

func main(){
	InitLogger()
    defer zlog.Sync()
	zlog.Infof("hello iponline")
	for(true){

	}
}



