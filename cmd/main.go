package cmd

import (
	"fmt"
	"os"

	"github.com/oswaldom-code/load-data-stock-donation-project/pkg/config"
	"github.com/oswaldom-code/load-data-stock-donation-project/src/adapters/cli"

	"github.com/spf13/cobra"
)

func init() {
	config.Load()
}

func newCliCmd() *cobra.Command {
	var cliCmd = &cobra.Command{
		Use: "cli",
		//Args: cobra.MinimumNArgs(1),
		Short: "Command line interface for load data  application",
		Long:  "Command line interface for load data on stock donation project application",
		Run: func(cmd *cobra.Command, args []string) {
			err := cli.RunCliCmd(cmd, args)
			if err != nil {
				fmt.Println("Error:", err)
			}
		},
	}
	cliCmd.Flags().StringP("function", "", "", "Function to execute")
	cliCmd.Flags().StringP("target", "t", "", "Target to execute")
	cliCmd.Flags().StringP("path", "", "", "path to load data")
	return cliCmd
}

func Execute() {
	// TODO: get PROJECT_NAME from config
	var rootCmd = &cobra.Command{Use: "load-data"}
	// Todo: add subcommands import from sub commands
	rootCmd.AddCommand(newCliCmd())
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: '%s'", err)
		os.Exit(1)
	}
}
