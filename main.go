package main


import (
	"github.com/xiaofeng-zerone/iponline/handler"
)



func main() {
	handler.InitLogger()

	handler.LoadConf("./conf/app.yml")

	for(true){

	}
	handler.UnInitLogger()
}



