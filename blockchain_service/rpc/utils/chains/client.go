package chains

import (
	"context"
	"github.com/JopenChen/star_guard/blockchain_service/rpc/internal/config"
	"github.com/JopenChen/star_guard/blockchain_service/rpc/utils/chains/web3storage"
	"io/fs"
	"os"
)

type Chains struct {
	Client Chain
}

type Chain interface {
	InitClient(ctx context.Context, config config.Config) (err error)
	Put(context context.Context, file *os.File) (cid string, err error)
	Get(context context.Context, cidString string) (file fs.File, err error)
}

var _ Chain = &web3storage.W3sClient{}

func (c *Chains) New(ctx context.Context, config config.Config) (err error) {
	switch config.Chains.Type {
	case "w3s":
		c.Client = &web3storage.W3sClient{}
	}

	return
}
