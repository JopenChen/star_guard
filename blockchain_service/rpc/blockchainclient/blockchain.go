// Code generated by goctl. DO NOT EDIT.
// Source: blockchain_service.proto

package blockchainclient

import (
	"context"

	"github.com/JopenChen/star_guard/blockchain_service/rpc/blockchain"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateDataRequest  = blockchain.CreateDataRequest
	CreateDataResponse = blockchain.CreateDataResponse
	RetrieveRequest    = blockchain.RetrieveRequest
	RetrieveResponse   = blockchain.RetrieveResponse

	Blockchain interface {
		Create(ctx context.Context, in *CreateDataRequest, opts ...grpc.CallOption) (*CreateDataResponse, error)
		Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*RetrieveResponse, error)
	}

	defaultBlockchain struct {
		cli zrpc.Client
	}
)

func NewBlockchain(cli zrpc.Client) Blockchain {
	return &defaultBlockchain{
		cli: cli,
	}
}

func (m *defaultBlockchain) Create(ctx context.Context, in *CreateDataRequest, opts ...grpc.CallOption) (*CreateDataResponse, error) {
	client := blockchain.NewBlockchainClient(m.cli.Conn())
	return client.Create(ctx, in, opts...)
}

func (m *defaultBlockchain) Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*RetrieveResponse, error) {
	client := blockchain.NewBlockchainClient(m.cli.Conn())
	return client.Retrieve(ctx, in, opts...)
}