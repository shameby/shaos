package conf

import (
	"flag"
)

var BaseConfig *Base

type Base struct {
	ServerName *string
	AppKey     string
	Port       *string
}

func init() {
	BaseConfig = &Base{
		ServerName: flag.String("n", "data_default", "data server name"),
		Port:       flag.String("p", "8081", "data grpc port"),
	}
	BaseConfig.AppKey = "FHSAIUH9Q82YE92HDABFSHJSB2"
	flag.Parse()
}
