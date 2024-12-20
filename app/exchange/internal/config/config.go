package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	MultiChainRpcConf struct {
		Endpoint string
	}
	SignRpcConf struct {
		Endpoint string
	}
	WalletAccountRpcConf struct {
		Endpoint string
	}
	NotifyRpcConf struct {
		Endpoint string
	}
	DbConf struct {
		Master DbConfig
		Slave  DbConfig
	}
}

type DbConfig struct {
	DbHost     string
	DbPort     string
	DbName     string
	DbUsername string
	DbPassword string
}
