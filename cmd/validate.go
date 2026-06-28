/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"
	"ue5cli/cmd/config"
	"ue5cli/cmd/discover"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// TODO: this is a discover operation - validate should check that the viper configuration file is valid
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates the UE5 project in the current folder, and the environment configuration.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("🤖 Validating project...")

		projectHome := viper.GetString("projectdir")
		projectPath, err := discover.FindUprojectInPath(projectHome)

		if err != nil {
			fmt.Printf("❌ %s\n", err.Error())
		} else {
			fmt.Printf("✅ Project path found at: %s\n", projectPath)
		}

		enginePath, err := discover.FindUe5EnginePath()
		if err != nil {
			fmt.Printf("❌ %s\n", err.Error())
		} else {
			viper.SetDefault("enginePath", enginePath)
			rootPath := strings.Replace(enginePath, "/Engine", "", 1)
			viper.SetDefault("rootPath", rootPath)
			fmt.Printf("✅ Engine path found at: %s\n", enginePath)
		}

		editorPath, err := discover.FindUnrealEditorPath(enginePath)
		if err != nil {
			fmt.Printf("❌ %s\n", err.Error())
		} else {
			viper.SetDefault("editorPath", editorPath)
			fmt.Printf("✅ Editor path found at: %s\n", editorPath)
		}

		buildScriptPath, err := discover.FindBuildScriptPath(enginePath)
		if err != nil {
			fmt.Printf("❌ %s\n", err.Error())
		} else {
			viper.SetDefault("buildScriptPath", buildScriptPath)
			fmt.Printf("✅ Build script path found at: %s\n", buildScriptPath)
		}

		buildToolPath, err := discover.FindBuildToolPath(enginePath)
		if err != nil {
			fmt.Printf("❌ %s\n", err.Error())
		} else {
			viper.SetDefault("buildToolPath", buildScriptPath)
			fmt.Printf("✅ Build tool path found at: %s\n", buildToolPath)
		}

		config.WriteConfigFile()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
