package base

import (
	"github.com/BurntSushi/toml"
)

var _config *tomlConfig

func GetConfig() *tomlConfig {
	return _config
}

// 此struct存放各标题字段
type tomlConfig struct {
	Server Server
	Log    Log
	DB     DB
	MQs    map[string]MQ
	Redis  Redis
	Mail   Mail
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  int64
	WriteTimeout int64
}

type Log struct {
	Path  string
	Level string
}

type DB struct {
	Database    string
	Username    string
	Password    string
	Host        string
	MaxOpenConn int
	MaxIdleConn int
}

type Redis struct {
	MaxIdle  int
	Timeout  int
	Ip       string
	Port     string
	Password string
}

type Mail struct {
	Username string
	Password string
	Nickname string
	MailType string
}

type MQ struct {
	Ip           string
	Port         string
	Username     string
	Password     string
	Host         string
	Exchange     string
	Key          string
	Queue        string
	ChanRangeNum int
}

func ConfigInit() {
	c := new(tomlConfig)
	if _, err := toml.DecodeFile("conf/config.toml", &c); err != nil {
		panic(err)
	}
	_config = c
}
