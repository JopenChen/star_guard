package svc

import "C"
import (
	"context"
	"github.com/JopenChen/star_guard/blockchain_service/rpc/internal/config"
	"github.com/JopenChen/star_guard/blockchain_service/rpc/utils/chains"
)

type ServiceContext struct {
	Config           config.Config
	BlockchainClient chains.Chain
}

func NewServiceContext(c config.Config) *ServiceContext {

	chainsObj := &chains.Chains{}
	err := chainsObj.New(context.Background(), c)
	if err != nil {
		return nil
	}

	err = chainsObj.Client.InitClient(context.Background(), c)
	if err != nil {
		return nil
	}

	return &ServiceContext{
		Config:           c,
		BlockchainClient: chainsObj.Client,
	}
}
