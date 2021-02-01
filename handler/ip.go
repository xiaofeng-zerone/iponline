package handler

import (
	"net"
	"fmt"
	"strconv"
	"github.com/xiaofeng-zerone/iponline/utils"
)

const IP_CMD = "/sbin/ip"

//NicMaskPrefixLenGet
func NicMaskPrefixLenGet(mask string) int {
	var mask_prefix int
	net_mask := net.ParseIP(mask)
	if net_mask != nil {
		net_ipv4mask := net_mask.To4()
		ipv4mask := net.IPv4Mask(net_ipv4mask[0], net_ipv4mask[1], net_ipv4mask[2], net_ipv4mask[3])
		mask_prefix, _ = ipv4mask.Size()
	} else {
		mask_prefix, _ = strconv.Atoi(mask)
	}
	return mask_prefix
}

//NicAddrInfoAdd 
func NicAddrInfoAdd(intf_name string, ip string, mask string) {
	mask_prefix := NicMaskPrefixLenGet(mask)
	ipmask := fmt.Sprintf("%s/%d", ip, mask_prefix)
	//ip/mask set
	var ipcfg_param = []string{"address", "add", ipmask, "dev", intf_name,}
	utils.ShellCmdExec(IP_CMD, ipcfg_param)
	//interface up
	var intfup_param = []string{"link", "set", "dev", intf_name, "up",}
	utils.ShellCmdExec(IP_CMD, intfup_param)
}

//NicAddrAllDel
func NicAddrAllDel(intf_name string) {
	//clear all ip
	var ipv4clear_param = []string{"-4", "address", "flush", "dev", intf_name, "scope", "global",}
	utils.ShellCmdExec(IP_CMD, ipv4clear_param)
	var ipv6clear_param = []string{"-6", "address", "flush", "dev", intf_name, "scope", "global",}
	utils.ShellCmdExec(IP_CMD, ipv6clear_param)
	//interface up
	var intfdown_param = []string{"link", "set", "dev", intf_name, "up",}
	utils.ShellCmdExec(IP_CMD, intfdown_param)
}

//NicIpSet 
func NicIpSet() {
	gcfg := GetConf()
	nicCfg := gcfg.NicCfgs
	if len(nicCfg.DEVICE) > 0 &&
		len(nicCfg.IPADDR) > 0 && 
		len(nicCfg.NETMASK) > 0 {
		NicAddrAllDel(nicCfg.DEVICE)
		NicAddrInfoAdd(nicCfg.DEVICE, nicCfg.IPADDR, nicCfg.NETMASK)
	}
}


