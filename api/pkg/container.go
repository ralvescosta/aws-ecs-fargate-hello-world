package pkg

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ralvescosta/gokit/configs"
	configsBuilder "github.com/ralvescosta/gokit/configs_builder"
	"github.com/ralvescosta/gokit/logging"
	"go.uber.org/dig"
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

	return container, nil
}

func ProvideSignal() chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	return sig
}
