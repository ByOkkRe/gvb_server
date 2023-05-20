package config

import "fmt"

type Mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	//日志等级
	LogLevel     string `yaml:"log_level"`
	Config       string `yaml:"config"`
	MaxIdleConns int    `yaml:"max-idle-conns"`
	MaxOpenConns int    `yaml:"max-open-conns"`
	DB           string `yaml:"db"`
}

func (m Mysql) Dsn() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DB,
		m.Config,
	)
}
