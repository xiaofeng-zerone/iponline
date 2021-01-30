package main


import (
	"github.com/xiaofeng-zerone/iponlin/handler"
)



func main() {
	InitLogger()
    defer zlog.Sync()
	zlog.Infof("hello iponline")

	handler.LoadConf("./conf/app.conf")

	for(true){

	}
}



