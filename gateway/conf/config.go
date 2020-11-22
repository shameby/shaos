package conf

import (
	"flag"
)

var BaseConfig *Base

type Base struct {
	ServerName *string
	Port       *string
}

func init() {
	BaseConfig = &Base{
		ServerName: flag.String("n", "gateway_default", "gateway name"),
		Port: flag.String("p", "8080", "gateway http port"),
	}
	flag.Parse()
}