package logic

import (
	"context"
	"crypto/md5"
	"fmt"
	"os"

	"github.com/JopenChen/star_guard/blockchain_service/rpc/blockchain"
	"github.com/JopenChen/star_guard/blockchain_service/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *blockchain.CreateDataRequest) (resp *blockchain.CreateDataResponse, err error) {
	resp = new(blockchain.CreateDataResponse)

	md5Byte := md5.Sum(in.GetFileStream())

	file, err := os.Create(fmt.Sprintf("%x", md5Byte))
	if err != nil {
		return
	}
	defer file.Close()

	_, err = file.Write(in.FileStream)
	if err != nil {
		return
	}

	cidString, err := l.svcCtx.BlockchainClient.Put(context.Background(), file)
	if err != nil {
		return
	}

	resp.Cid = cidString
	resp.Msg = "OK"

	return
}
