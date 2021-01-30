package handler

import (
	"fmt"
	"time"
	"github.com/xiaofeng-zerone/iponline/utils"
)


//PingIp ping ip
func PingBaidu() bool {
	cmd := fmt.Sprintf("ping %s -c 2 > /dev/null 2>&1", "www.baidu.com")
	_, err := utils.ShellCmdExecBySh(cmd)
	return err
}



//PingCheck
func PingCheck() bool {
	var isOk bool
	for i:=0; i<10; i++ {
		if PingBaidu() {
			isOk = true
			break
		}
		fmt.Println("ping www.baidu.com is failed. count:", i)
		time.Sleep(5*time.Second)
	}
	
	return isOk
}