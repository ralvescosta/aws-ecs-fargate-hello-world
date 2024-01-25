package main

import (
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/api/cmd"
	"github.com/spf13/cobra"

	_ "github.com/ralvescosta/aws-ecs-fargate-hello-world/api/docs"
)

// rootCmd represents the base command when called without any subcommands
var root = &cobra.Command{
	Use:     "app",
	Short:   "HTTP API",
	Version: "0.0.1",
}

// @title Hello World HTTP API
// @version 1.0
// @description This is a sample server.
// @termsOfService https://github.com/ralvescosta/aws-ecs-fargate-hello-world/blob/main/LICENSE

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://github.com/ralvescosta/aws-ecs-fargate-hello-world/blob/main/LICENSE

// @host localhost:3333
// @BasePath /v1
func main() {
	root.AddCommand(cmd.ApiCmd)

	root.Execute()
}
