package main


import (
	"time"
	"github.com/xiaofeng-zerone/iponline/handler"
)



func main() {
	handler.InitLogger()
	handler.LoadConf("./conf/app.yml")

	time.Sleep(600*time.Second)
	for(true){
		
		//ping check
		if handler.PingCheck() == false {
			//dns set
			handler.DnsSet()
			//ip set
			handler.NicIpSet()
			//default route set
			handler.NicDefaultRouteSet()
		}

		time.Sleep(5*time.Second)
	}
	handler.UnInitLogger()
}



