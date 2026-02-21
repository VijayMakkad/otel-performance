package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "otel-benchmark",
	Short: "otel-benchmark is a CLI tool for benchmarking OpenTelemetry collector",
	Long:  `otel-benchmark is a CLI tool for benchmarking OpenTelemetry collector`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of otel-benchmark",
	Long:  `Print the version number of otel-benchmark`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("otel-benchmark version 1.0.0")
	},
}
