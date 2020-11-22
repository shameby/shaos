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
		ServerName: flag.String("n", "meta_default", "meta server name"),
		Port: flag.String("p", "8083", "meta server grpc port"),
	}
	flag.Parse()
}