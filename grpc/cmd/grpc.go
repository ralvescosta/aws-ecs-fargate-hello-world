package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/ralvescosta/ec2-hello-world/protos"
)

type GRPCParams struct {
	CommonParams

	Sig            chan os.Signal
	ProductsServer protos.ProductsServer
}

func gRPC(params GRPCParams) error {
	params.Logger.Debug("Stating gRPC Server...")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", params.Cfg.HTTPConfigs.Port))
	if err != nil {
		params.Logger.Error("failed to create tcp listener", zap.Error(err))
		return err
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	protos.RegisterProductsServer(grpcServer, params.ProductsServer)

	params.Logger.Debug("gRPC Server started!")

	return grpcServer.Serve(lis)
}

var GRPCCmd = &cobra.Command{
	Use:   "grpc",
	Short: "gRPC Server Command",
	RunE:  RunCommand(gRPC),
}
