package handler

var (
	config *Conf
)

// NicInterfacesOpt  
type NicInterfacesOpt struct {
	ONBOOT           string `yaml:"ONBOOT"`
	
}

//Conf Conf
type Conf struct {
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

	err := yaml.Unmarshal(cfg, &config)
	if err != nil {
		zlog.Errorf("yaml Unmarshal error: %v", err)
		return err
	}

	zlog.Infof("config dump:", config)
}

//Get return config
func Get() *Conf {
	return config
}



