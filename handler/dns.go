package handler

import (
	"fmt"
	"io/ioutil"
	"github.com/xiaofeng-zerone/iponline/utils"
)

//DnsCheck
func DnsCheck() {
	var dnscfg string

	gcfg := GetConf()
	if len(gcfg.DnsDefault) == 0 {
		return
	}

	cfg, err := ioutil.ReadFile("/etc/resolv.conf")
	if err != nil {
		zlog.Errorf("read file failed.", err)
	} else {
		dnscfg = string(cfg)
	}

	if len(dnscfg) == 0 {
		//dns set 
		cmd := fmt.Sprintf("echo \"nameserver %s\" >> /etc/resolv.conf", gcfg.DnsDefault)
		utils.ShellCmdExecBySh(cmd)
	}
	
	return
}