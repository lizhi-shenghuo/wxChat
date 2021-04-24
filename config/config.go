package config

import (
	"flag"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var (
	cfgFile = flag.String("config", "./config.yaml", "配置文件路径")

	cfg *AppConfig
)

//Config example config
type AppConfig struct {
	Listen                 string `yaml:"listen"`
	*Redis                 `yaml:"redis"`
	*OfficialAccountConfig `yaml:"officialAccountConfig"`
	*Mysql                 `yaml:"mysql"`
}

type Redis struct {
	Host        string `yaml:"host"`
	Password    string `yaml:"password"`
	Database    int    `yaml:"database"`
	MaxActive   int    `yaml:"maxActive"`
	MaxIdle     int    `yaml:"maxIdle"`
	IdleTimeout int    `yaml:"idleTimeout"`
}

//OfficialAccountConfig 公众号相关配置
type OfficialAccountConfig struct {
	AppID          string `yaml:"appID"`
	AppSecret      string `yaml:"appSecret"`
	Token          string `yaml:"token"`
	EncodingAESKey string `yaml:"encodingAESKey"`
}

type Mysql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"dbName"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	LogZap       string `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
}

//GetConfig 获取配置
func GetConfig() *AppConfig {
	if cfg != nil {
		return cfg
	}
	bytes, err := ioutil.ReadFile(*cfgFile)
	if err != nil {
		panic(err)
	}

	cfgData := &AppConfig{}
	err = yaml.Unmarshal(bytes, cfgData)
	if err != nil {
		panic(err)
	}
	cfg = cfgData
	return cfg
}
