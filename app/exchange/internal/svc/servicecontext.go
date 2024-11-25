package svc

import (
	"business-platform/app/common/grpcclient/multichain"
	"business-platform/app/exchange/internal/config"
)

type ServiceContext struct {
	Config           config.Config
	MultichainClient multichain.GrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	multichainClient, err := multichain.NewGrpcClient(c.MultichainRpcConf.Endpoint)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:           c,
		MultichainClient: multichainClient,
	}
}
