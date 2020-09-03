package conf

type Config struct {
	Mysql struct {
		User string `yaml:"user"`
		Host string `yaml:"host"`
		Password string `yaml:"password"`
		Port int `yaml:"port"`
		DbName string `yaml:"dbname"`
		ConnetMax int `yaml:"connetMax"`
	}
}
