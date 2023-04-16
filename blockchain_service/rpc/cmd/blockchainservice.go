package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/JopenChen/star_guard/blockchain_service/rpc/blockchain"
	"github.com/JopenChen/star_guard/blockchain_service/rpc/internal/config"
	"github.com/JopenChen/star_guard/blockchain_service/rpc/internal/server"
	"github.com/JopenChen/star_guard/blockchain_service/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func BlockchainService(cmd *cobra.Command, args []string) {
	var c config.Config
	conf.MustLoad(args[2], &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		blockchain.RegisterBlockchainServer(grpcServer, server.NewBlockchainServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
