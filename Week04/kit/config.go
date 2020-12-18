package kit

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// DBConfig db config
type DBConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
	Charset  string `yaml:"charset"`
}

// Config config
type Config struct {
	Mysql DBConfig `yaml:"mysql"`
}

// LoadConfig 加载配置文件
func LoadConfig(conf *Config) error {
	path := "./config/app.yml"
	appConfig, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(appConfig, conf)
}
