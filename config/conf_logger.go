package config

type Logger struct {
	//日志显示级别
	Level string `yaml:"level"`
	//日志前缀
	Prefix string `yaml:"prefix"`
	//日志文件
	Director string `yaml:"director"`
	//行号
	RowNumber bool `yaml:"row_number"`
	//打印路径
	LogConsole bool `yaml:"log_console"`
}
