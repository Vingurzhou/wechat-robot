package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	DSN  string
	GEWE struct {
		Token string
		AppId string
	}
}
