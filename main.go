package main


import (
	"time"
	"github.com/xiaofeng-zerone/iponline/handler"
)



func main() {
	handler.InitLogger()

	handler.LoadConf("./conf/app.yml")

	for(true){
		handler.DnsCheck()
		handler.PingCheck()

		time.Sleep(5*time.Second)
	}
	handler.UnInitLogger()
}



