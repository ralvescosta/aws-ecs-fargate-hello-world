package cmd

import (
	"go.uber.org/dig"
	"go.uber.org/zap"

	"github.com/ralvescosta/gokit/configs"
	"github.com/ralvescosta/gokit/logging"
	"github.com/spf13/cobra"

	"github.com/ralvescosta/aws-ecs-fargate-hello-world/api/pkg"
)

type CommonParams struct {
	dig.In

	Cfg    *configs.Configs
	Logger logging.Logger
}

func RunCommand(runner any) func(*cobra.Command, []string) error {
	return func(c *cobra.Command, s []string) error {
		ioc, err := pkg.NewContainer()
		if err != nil {
			return err
		}

		err = ioc.Invoke(func(logger logging.Logger) error {
			if err := ioc.Invoke(runner); err != nil {
				logger.Error("error running command", zap.Error(err))
				return err
			}

			return nil
		})

		return err
	}
}
