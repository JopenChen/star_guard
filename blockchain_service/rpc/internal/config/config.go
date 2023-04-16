package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Chains Chains
}

type Chains struct {
	Type              string
	Web3StorageConfig Web3StorageConfig
}

type Web3StorageConfig struct {
	Endpoint string `json:",optional"`
	Token    string
}
