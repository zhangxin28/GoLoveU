package config

import (
	"goloveu/utils"
)

// Conf represents the bbs-go.yaml object
var Conf *Config

// Config represents the bbs-go.yaml object type
type Config struct {
	Env        string `yaml:"Env"`        // 环境：prod、dev
	BaseURL    string `yaml:"BaseURL"`    // base url
	Port       string `yaml:"Port"`       // 端口
	LogFile    string `yaml:"LogFile"`    // 日志文件
	ShowSQL    bool   `yaml:"ShowSQL"`    // 是否显示日志
	StaticPath string `yaml:"StaticPath"` // 静态文件目录

	MySQLURL string `yaml:"MySQLURL"` // 数据库连接地址

	// Github
	Github struct {
		ClientID     string `yaml:"ClientID"`
		ClientSecret string `yaml:"ClientSecret"`
	} `yaml:"Github"`

	// QQ登录
	QQConnect struct {
		AppID  string `yaml:"AppID"`
		AppKey string `yaml:"AppKey"`
	} `yaml:"QQConnect"`

	// 阿里云oss配置
	AliyunOss struct {
		Host         string `yaml:"Host"`
		Bucket       string `yaml:"Bucket"`
		Endpoint     string `yaml:"Endpoint"`
		AccessID     string `yaml:"AccessID"`
		AccessSecret string `yaml:"AccessSecret"`
	} `yaml:"AliyunOss"`

	// 百度ai
	BaiduAi struct {
		APIKey    string `yaml:"APIKey"`
		SecretKey string `yaml:"SecretKey"`
	} `yaml:"BaiduAi"`

	// 百度SEO相关配置
	// 文档：https://ziyuan.baidu.com/college/courseinfo?id=267&page=2#h2_article_title14
	BaiduSEO struct {
		Site  string `yaml:"Site"`
		Token string `yaml:"Token"`
	} `yaml:"BaiduSEO"`

	// smtp
	SMTP struct {
		Host     string `yaml:"Host"`
		Port     int `yaml:"Port"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
		SSL      bool   `yaml:"SSL"`
	} `yaml:"SMTP"`
}

// InitConfig performs the bbs-go.yaml configuration init
func InitConfig(filename string) {
	Conf = &Config{}
	err := utils.YamlUnmarshal(filename, Conf)
	if err != nil {
		utils.LogError(err)
	}
}
