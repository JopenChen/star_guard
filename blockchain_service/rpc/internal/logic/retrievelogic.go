package logic

import (
	"context"
	"fmt"
	"io"

	"github.com/JopenChen/star_guard/blockchain_service/rpc/blockchain"
	"github.com/JopenChen/star_guard/blockchain_service/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveLogic {
	return &RetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RetrieveLogic) Retrieve(in *blockchain.RetrieveRequest) (resp *blockchain.RetrieveResponse, err error) {
	resp = new(blockchain.RetrieveResponse)
	resp.FileSteam = []byte{}

	file, err := l.svcCtx.BlockchainClient.Get(context.Background(), in.GetCid())
	if err != nil {
		return
	}

	for {
		n, err1 := file.Read(resp.FileSteam)
		fmt.Println(n)
		if err1 != nil {
			if err1 == io.EOF {
				break
			}
			return nil, err1
		}
	}

	resp.Msg = "OK"

	return
}
