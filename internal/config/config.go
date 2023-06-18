package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string
	Cache      cache.CacheConf
	Jwt        struct {
		AccessSecret  string
		AccessPublic  string
		AccessExpire  int64
		RefreshExpire int64
	} `json:"jwt"`
	Mail struct {
		Login    string
		Identity string
		From     string
		Password string
		SmtpHost string
		SmtpPort string
	} `json:"mail"`
}
