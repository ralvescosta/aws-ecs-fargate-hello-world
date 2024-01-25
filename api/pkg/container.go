package pkg

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ralvescosta/gokit/configs"
	"github.com/ralvescosta/gokit/logging"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/ralvescosta/aws-ecs-fargate-hello-world/api/pkg/handlers"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/api/pkg/internal/services"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/protos"
	configsBuilder "github.com/ralvescosta/gokit/configs_builder"
)

func NewContainer() (*dig.Container, error) {
	cfg, err := configsBuilder.NewConfigsBuilder().
		HTTP().
		Build()

	if err != nil {
		return nil, err
	}

	container := dig.New()

	container.Provide(func() *configs.Configs { return cfg })
	container.Provide(func() *configs.AppConfigs { return cfg.AppConfigs })
	container.Provide(logging.NewDefaultLogger)
	container.Provide(ProvideSignal)
	container.Provide(handlers.NewProductsHandler)
	container.Provide(ProvideProductsClient)
	container.Provide(services.NewProductsService)

	return container, nil
}

func ProvideSignal() chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	return sig
}

func ProvideProductsClient(cfgs *configs.Configs, logger logging.Logger) (protos.ProductsClient, error) {
	logger.Debug("connection to products grpc...")

	conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("failure to stablish connection", zap.Error(err))
		return nil, err
	}

	logger.Debug("products grpc connected!")

	return protos.NewProductsClient(conn), nil
}
