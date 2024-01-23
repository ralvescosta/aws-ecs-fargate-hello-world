package cmd

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"

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
	params.Logger.Debug("creating tpc listener...")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", params.Cfg.HTTPConfigs.Port))
	if err != nil {
		params.Logger.Error("failed to create tcp listener", zap.Error(err))
		return err
	}

	params.Logger.Debug("tpc listener created!")

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	protos.RegisterProductsServer(grpcServer, params.ProductsServer)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		<-params.Sig
		params.Logger.Info("got signal, attempting graceful shutdown...")
		grpcServer.GracefulStop()
		wg.Done()
	}()

	params.Logger.Debug("starting gRPC server...")

	go func() {
		time.Sleep(time.Second)
		params.Logger.Debug("gRPC server started!")
	}()

	if err := grpcServer.Serve(lis); err != nil {
		params.Logger.Error("failed to started gRPC server", zap.Error(err))
		return err
	}

	wg.Wait()
	params.Logger.Info("clean shutdown")

	return nil
}

var GRPCCmd = &cobra.Command{
	Use:   "grpc",
	Short: "gRPC Server Command",
	RunE:  RunCommand(gRPC),
}
