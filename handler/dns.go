package handler

import (
	"fmt"
	"strings"
	"io/ioutil"
	"github.com/xiaofeng-zerone/iponline/utils"
)

//DnsSet
func DnsSet() bool {
	var dnscfg string

	gcfg := GetConf()
	if len(gcfg.DnsDefault) == 0 {
		return false
	}

	cfg, err := ioutil.ReadFile("/etc/resolv.conf")
	if err != nil {
		zlog.Errorf("read file failed.", err)
	} else {
		dnscfg = string(cfg)
	}

	if len(dnscfg) == 0 || strings.Contains(dnscfg, gcfg.DnsDefault) == false {
		//dns set 
		cmd := fmt.Sprintf("echo \"nameserver %s\" >> /etc/resolv.conf", gcfg.DnsDefault)
		utils.ShellCmdExecBySh(cmd)
	}
	
	return true
}