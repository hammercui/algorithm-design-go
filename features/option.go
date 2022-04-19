package features

import "log"

/**
可选参数使用
 */


type Config struct {
	IP string
	Port int32
	Pass string
	UserName string
}

type Options func(config *Config)

func WithIP(ip string) Options  {
	return func(config *Config) {
		config.IP = ip
	}
}

func WithPort(port int32) Options  {
	return func(config *Config) {
		config.Port = port
	}
}

func DefaultConfig(conf *Config) *Config {
	conf.Pass  = "Pass"
	conf.UserName = "hammer"
	return conf
}

func NewConn(options ...Options) *Config  {
	conf := &Config{}
	//默认值设定
	conf = DefaultConfig(conf)
	//遍历可选参数
	for _,op := range options{
		op(conf)
	}
	return conf
}

func main()  {
	//传入自定义可选参数
	var options = []Options{
		WithIP("127.0.0.1"),
		WithPort(8080),
	}
	//后续有新配置，可以持续更新，而不会破坏原有结构
	conf := NewConn(options...)
	log.Println("conf",conf)
}
