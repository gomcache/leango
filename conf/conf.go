package conf

import (
	"LearnGoProject/utils"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

var config *Config
var once sync.Once

func Settings() *Config  {
	once.Do(func() {
		loadLocalConfig()
	})
	return config
}
func loadLocalConfig()  {

	config = new(Config)

	yamlFile, err := ioutil.ReadFile("conf/config.yml")
	if err != nil {
		utils.ERR.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	// err = yaml.Unmarshal(yamlFile, &resultMap)
	if err != nil {
		utils.ERR.Fatalf("Unmarshal: %v", err)
	}
	utils.System.Printf("配置加载成功: %v", config.Mysql)
}
