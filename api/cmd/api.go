package cmd

import (
	"os"

	"github.com/ralvescosta/gokit/httpw/server"
	"github.com/spf13/cobra"

	"github.com/ralvescosta/aws-ecs-fargate-hello-world/api/pkg/handlers"
)

type APIParams struct {
	CommonParams

	Sig      chan os.Signal
	Handlers handlers.HTTPHandlers
}

func api(params APIParams) error {
	params.Logger.Debug("Stating HTTP API...")

	router := server.
		NewHTTPServerBuilder(params.Cfg.HTTPConfigs, params.Logger).
		Signal(params.Sig).
		WithOpenAPI().
		Build()

	params.Handlers.Install(router)

	return router.Run()
}

var ApiCmd = &cobra.Command{
	Use:   "api",
	Short: "API Server Command",
	RunE:  RunCommand(api),
}
