package sf

import (
	"github.com/storezhang/gox"
)

var (
	sf *gox.Snowflake
)

func Init() {
	var err error

	if sf, err = gox.NewSnowflake(0); nil != err {
		panic("init sf fail")
	}
}

func Id() int64 {
	return sf.NextId()
}

func Ids() string {
	return sf.Next().String()
}
