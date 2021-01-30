package handler

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

var (
	config *Conf
)

// NicInterfacesOpt  
type NicInterfacesOpt struct {
	DEVICE				string `yaml:"DEVICE"`
	IPADDR          	string `yaml:"IPADDR"`
	NETMASK           	string `yaml:"NETMASK"`
}

//Conf Conf
type Conf struct {
	DnsDefault string `yaml:"DnsDefault"`
	Ipv4GatewayDefault string `yaml:"Ipv4GatewayDefault"`
	Ipv6GatewayDefault string `yaml:"Ipv6GatewayDefault"`
	NicCfgs NicInterfacesOpt `yaml:"NicCfgs"`

}

// LoadConf
func LoadConf(path string) error {
	config = &Conf{}

	cfg, err := ioutil.ReadFile(path)
	if err != nil {
		zlog.Errorf("read file failed.", err)
		return err
	}

	err = yaml.Unmarshal(cfg, &config)
	if err != nil {
		zlog.Errorf("yaml Unmarshal error: %v", err)
		return err
	}

	zlog.Infof("config dump:", config)
	return nil
}

//Get return config
func GetConf() *Conf {
	return config
}



