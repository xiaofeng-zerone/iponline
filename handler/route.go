package handler

import (
	"github.com/xiaofeng-zerone/iponline/utils"
)

const ROUTE_CMD = "/sbin/route"
const DEFAULT_METRIC = "0"

//NicDefaultRouteCfgClear
func NicDefaultRouteCfgClear() bool {
	var ipv4route_del_param = []string{"del", "default"}
	cmdret, ret := utils.ShellCmdExec(ROUTE_CMD, ipv4route_del_param)
	if len(cmdret) > 0 || ret == false {
		return false
	}
	return true
}

//NicIpv4DefaultRouteCfgAdd
func NicIpv4DefaultRouteCfgAdd(gateway string) bool {
	if len(gateway) == 0 {
		return false
	}
	var ipv4route_add_param = []string{"-A", "inet", "add", "default", "gw", gateway, "metric", DEFAULT_METRIC,}
	cmdret, ret := utils.ShellCmdExec(ROUTE_CMD, ipv4route_add_param)
	if len(cmdret) > 0 || ret == false {
		return false
	}
	return true
}


//NicIpv6DefaultRouteCfgAdd
func NicIpv6DefaultRouteCfgAdd(gateway string) bool {
	if len(gateway) == 0 {
		return false
	}
	var ipv6route_add_param = []string{"-A", "inet6", "add", "default", "gw", gateway, "metric", DEFAULT_METRIC,}
	cmdret, ret := utils.ShellCmdExec(ROUTE_CMD, ipv6route_add_param)
	if len(cmdret) > 0 || ret == false {
		return false
	}
	return true
}

//NicDefaultRouteSet 
func NicDefaultRouteSet() {
	gcfg := GetConf()

	NicDefaultRouteCfgClear()
	NicIpv4DefaultRouteCfgAdd(gcfg.Ipv4GatewayDefault)
	NicIpv6DefaultRouteCfgAdd(gcfg.Ipv6GatewayDefault)
}



