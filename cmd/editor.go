/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"ue5cli/cmd/config"
)

var editorCmd = &cobra.Command{
	Use:   "editor",
	Short: "Run the Unreal Editor.",
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("🚀 Running Unreal Editor...")
		editorCmd := exec.Command(config.Ue5EditorPath, config.Ue5ProjectFilePath)

		err := editorCmd.Start()

		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(editorCmd)
}
