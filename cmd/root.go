/*
Copyright © 2026 John-Evans Wagenaar <jewagenaar@gmail.com>
*/
package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"os"
	"ue5cli/cmd/config"
)

var rootCmd = &cobra.Command{
	Use:   "ue5cli",
	Short: "Unreal Engine 5 CLI tool for macOS.",
}

func ExecuteContext(ctx context.Context) {
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	config.InitConfig()
}
