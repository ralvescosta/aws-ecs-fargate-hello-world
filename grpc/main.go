package main

import (
	"github.com/ralvescosta/ecs-hello-world/grpc/cmd"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var root = &cobra.Command{
	Use:     "app",
	Short:   "HTTP API",
	Version: "0.0.1",
}

func main() {
	root.AddCommand(cmd.GRPCCmd)

	root.Execute()
}
